package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func mainPage_(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\n", req.Method)

	body += "Header ===============\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\n", k, v)
	}

	body += "Query parameters ===============\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\n", k, v)
	}

	log.Println("req.Header.User-Agent:", req.Header["User-Agent"])
	log.Println("req.URL.name:", req.URL.Query().Get("name"))

	res.Write([]byte(body))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)

	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	body += "Query parameters ===============\r\n"
	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	log.Println("req.Header.User-Agent:", req.Header["User-Agent"])
	log.Println("req.Form.Get(\"name\"):", req.Form.Get("name"))
	log.Println("req.FormValue(\"name\"):", req.FormValue("name"))

	res.Write([]byte(body))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	res.Write([]byte("Это страница /api"))
}

func postsPage(res http.ResponseWriter, req *http.Request) {
	log.Println("Method", req.Method)

	if req.Method != http.MethodPost {
		http.Error(res, "Only POST___ requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	type Subj struct {
		Product string `json:"name"`
		Price   int    `json:"price"`
	}

	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	res.Header().Set("content-type", "application/json")
	res.Header().Set("x-my-header", "my-own-header")
	// устанавливаем код 200
	res.WriteHeader(http.StatusCreated)
	// пишем тело ответа
	res.Write(resp)

}

func postsPageUrlencoded(res http.ResponseWriter, req *http.Request) {
	log.Println("Method", req.Method)

	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	////////////////// for application/x-www-form-urlencoded /////////////

	log.Println("Received name:", req.FormValue("name"))
	log.Println("Received age:", req.FormValue("age"))

	err := req.ParseForm()
	if err != nil {
		log.Println("Ошибка при req.ParseForm()")
	}
	// Получение всех значений параметра "name" из формы
	names := req.Form["name"]
	// Получение всех значений параметра "age" из формы
	ages := req.Form["age"]

	// Вывод всех значений в лог
	for _, name := range names {
		log.Println("Received name:", name)
	}
	for _, age := range ages {
		log.Println("Received age:", age)
	}

	res.Write([]byte("Это страница /Posts"))
}

func postsPageJson(res http.ResponseWriter, req *http.Request) {
	log.Println("Method", req.Method)

	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	//////////////////////// for application/json ///////////////////////
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Ошибка при io.ReadAll")
	}
	log.Println("req.Body:", req.Body)
	log.Println("body:", body)

	type MyData struct {
		Name string
		Age  int
	}

	var data MyData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Ошибка при разборе JSON:", err)
	}

	log.Println("Received name:", data.Name)
	log.Println("Received age:", data.Age)

	res.Write([]byte("Это страница /Posts"))
	io.WriteString(res, "1")
	fmt.Fprint(res, "2")
	res.Write([]byte("3"))
	// Это страница /Posts123
}

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/auth/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func authPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Добро пожаловать!")
		} else {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/posts/`, postsPage)
	mux.HandleFunc(`/posts/urlencoded`, postsPageUrlencoded)
	mux.HandleFunc(`/posts/json`, postsPageJson)
	mux.HandleFunc(`/auth/`, authPage)

	err := http.ListenAndServe(`:3000`, mux)
	if err != nil {
		panic(err)
	}
}

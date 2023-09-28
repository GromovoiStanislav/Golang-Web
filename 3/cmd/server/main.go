package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
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
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	//////////////////////// for application/json ///////////////////////
	//body, err := io.ReadAll(req.Body)
	//if err != nil {
	//	log.Println("Ошибка при io.ReadAll")
	//}
	//log.Println("req.Body:", req.Body)
	//log.Println("body:", body)
	//
	//type MyData struct {
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//}
	//
	//var data MyData
	//err = json.Unmarshal(body, &data)
	//if err != nil {
	//	log.Println("Ошибка при разборе JSON:", err)
	//}
	//
	//log.Println("Received name:", data.Name)
	//log.Println("Received age:", data.Age)

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/posts/`, postsPage)

	err := http.ListenAndServe(`:3000`, mux)
	if err != nil {
		panic(err)
	}
}

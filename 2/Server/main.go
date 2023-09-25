package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts      []Post
	postsMutex sync.RWMutex
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello, World!")
	w.Write([]byte("Hello World!"))
}

type Message struct {
	Text string `json:"message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	message := Message{Text: "Hello, JSON!"}

	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Ошибка при маршалинге JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPosts(w, r)
	case http.MethodPost:
		createPost(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	postsMutex.RLock()
	defer postsMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}

	var newPost Post
	err = json.Unmarshal(body, &newPost)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}

	postsMutex.Lock()
	defer postsMutex.Unlock()

	newPost.ID = len(posts) + 1
	posts = append(posts, newPost)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Создан новый пост с ID: %d", newPost.ID)
}

func getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	postsMutex.RLock()
	defer postsMutex.RUnlock()

	// Извлекаем значение id из параметра маршрута вручную
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, "Некорректный URL", http.StatusBadRequest)
		return
	}

	idStr := pathSegments[len(pathSegments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/hello", helloHandler) // обработчик для пути "/hello"
	http.HandleFunc("/json", jsonHandler)   // обработчик для пути "/json"

	http.HandleFunc("/post", postHandler)         // обработчик для пути "/post"
	http.HandleFunc("/post/", getPostByIDHandler) // Обратите внимание на обработчик /post/

	http.ListenAndServe(":8080", nil) // Запуск HTTP-сервера на порту 8080
	fmt.Println("Server is listening on :8080...")

}

// go run main.go

// go mod init mymodule
// go build -o myapp.exe
// ./myapp
// Ctrl+C

//go build main.go
// go build -o myapp.exe main.go
// ./myapp
// Ctrl+C

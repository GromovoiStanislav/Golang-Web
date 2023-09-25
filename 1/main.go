package main

import (
	"log"
	"net/http"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Web!"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", HelloWeb)
	http.HandleFunc("/world", HelloWorld)

	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Println("listen and serve:", err)
	}
}

//go run main.go
//go build main.go

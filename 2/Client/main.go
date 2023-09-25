package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	{
		resp, err := http.Get("http://localhost:8080/hello") // Отправляем GET-запрос на сервер
		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body) // Чтение ответа
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}

		fmt.Println("Ответ сервера:", string(body))
	}

	{
		resp, err := http.Get("http://localhost:8080/json") // Отправляем GET-запрос на сервер
		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body) // Чтение ответа
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}

		fmt.Println("Ответ сервера:", string(body))
	}

	{
		url := "http://localhost:8080/post" // URL сервера, куда отправляем POST-запрос

		// Тело POST-запроса в виде строки
		payload := []byte(`{"title": "Title 1", "content": "content..."}`)

		// Создаем новый POST-запрос
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Ошибка при создании запроса:", err)
			return
		}

		// Устанавливаем заголовок Content-Type
		req.Header.Set("Content-Type", "application/json")

		// Отправляем запрос на сервер
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			return
		}
		defer resp.Body.Close()

		// Читаем ответ сервера
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}

		fmt.Println("Ответ сервера:", string(body))
		// Прочитать статус ответа
		fmt.Println("Статус ответа:", resp.Status)   // Это строка
		fmt.Println("Код статуса:", resp.StatusCode) // Это числовой код статуса
	}

	{
		// Создаем URL с URI-параметром (например, /post/3)
		url := "http://localhost:8080/post/1"

		// Выполняем GET-запрос
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body) // Чтение ответа
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}

		fmt.Println("Ответ сервера:", string(body))
	}

	{

		url := "http://localhost:8080/post"

		// Выполняем GET-запрос
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body) // Чтение ответа
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}

		fmt.Println("Ответ сервера:", string(body))
	}
}

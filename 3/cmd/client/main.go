package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendGET() {
	// URL для отправки GET запроса
	url := "http://localhost:3000/?id=12345&name=John%20Doe&filter=town&filter=country"

	// Отправка GET запроса
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Вывод ответа на экран
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}

func sendGET_api() {
	// URL для отправки GET запроса
	url := "http://localhost:3000/api"

	// Отправка GET запроса
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Вывод ответа на экран
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}

func sendPOST() {
	// URL для отправки POST запроса
	url := "http://localhost:3000/posts/"

	// Тело POST запроса (в данном случае, пустое тело)
	requestBody := []byte("")

	// Отправка POST запроса
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение статуса ответа
	fmt.Println("Статус ответа:", response.Status)

	// Чтение заголовков ответа
	fmt.Println("Заголовки ответа:")
	for key, values := range response.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Вывод ответа на экран
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}

func sendPOST_JSON() {
	// URL для отправки POST запроса
	url := "http://localhost:3000/posts/json"

	// Данные для тела POST запроса в формате JSON
	data := map[string]interface{}{
		"name": "Stas",
		"age":  30,
	}

	// Кодируем данные в JSON
	requestBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON:", err)
		return
	}

	// Отправка POST запроса
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Вывод ответа на экран
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}

func sendPOST_urlencoded() {
	// URL для отправки POST запроса
	url := "http://localhost:3000/posts/urlencoded"

	// Данные для тела POST запроса
	data := "name=Tom&age=30&name=Stas"

	// Отправка POST запроса с данными в формате x-www-form-urlencoded
	response, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Вывод ответа на экран
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}

func main() {
	sendGET()
	sendGET_api()
	sendPOST()
	sendPOST_JSON()
	sendPOST_urlencoded()
}

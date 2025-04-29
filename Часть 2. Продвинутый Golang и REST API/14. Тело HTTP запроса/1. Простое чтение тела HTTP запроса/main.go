package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("failed to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	fmt.Println("Полученное тело входящего HTTP запроса:", httpRequestBodyString)
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

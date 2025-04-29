package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, World!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err)
	} else {
		fmt.Println("Я корректно обработал HTTP запрос")
	}
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

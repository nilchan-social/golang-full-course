package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платёж обработан!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err)
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err)
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}

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
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

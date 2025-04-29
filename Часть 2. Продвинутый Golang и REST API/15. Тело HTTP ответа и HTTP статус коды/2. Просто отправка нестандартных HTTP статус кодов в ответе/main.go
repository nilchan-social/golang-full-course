package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	// w.WriteHeader(http.StatusConflict)
	w.WriteHeader(http.StatusInternalServerError)

	fmt.Println("Успешно записал статус код в HTTP ответ!")
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

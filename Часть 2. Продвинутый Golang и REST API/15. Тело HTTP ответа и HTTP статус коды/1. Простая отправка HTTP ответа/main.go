package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is my HTTP response!"))
	if err != nil {
		fmt.Println("Ошибка записи HTTP ответа:", err)
		return
	}

	fmt.Println("Успешно записал HTTP ответ!")
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

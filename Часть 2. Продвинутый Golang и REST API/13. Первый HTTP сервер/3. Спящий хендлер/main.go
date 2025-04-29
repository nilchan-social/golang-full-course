package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("Handler отработал успешно!")
}

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	_, err := w.Write([]byte("HTTP response!"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("SleepHandler отработал успешно!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sleep", sleepHandler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

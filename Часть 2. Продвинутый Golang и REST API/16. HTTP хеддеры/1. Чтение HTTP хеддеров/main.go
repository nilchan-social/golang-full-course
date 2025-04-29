package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("Название хеддера:", k, "-- значение, лежащее в хеддере:", v)
	}

	fmt.Println("payHandler закончил своё выполнение.")
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

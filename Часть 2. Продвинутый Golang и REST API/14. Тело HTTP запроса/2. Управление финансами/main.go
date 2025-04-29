package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0     // usd

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("failed to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("failed to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount
		fmt.Println("Оплата прошла успешно:", money)
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("failed to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("failed to convert string to integer:", err)
		return
	}

	mtx.Lock()
	if saveAmount <= money {
		bank += saveAmount
		money -= saveAmount
		fmt.Println("Денег в копилке:", bank)
		fmt.Println("Балланс кошелька:", money)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}

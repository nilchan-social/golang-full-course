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
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount

		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
	}
	mtx.Unlock()
}

// Нет добавленных HTTP ответов и HTTP статус кодов
func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert string to integer:", err)
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

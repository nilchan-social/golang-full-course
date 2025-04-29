package main

import "fmt"

func main() {
	// Создали nil канал. nil потому что не проинициализирован через make
	var ch chan string
	// var ch chan string = make(chan string) // стоит проинициализировать через make, как сразу всё заработает

	go func() {
		// Запись в nil канал. Получаем блокировку.
		ch <- "hello"
	}()

	// Чтение из nil канала. Получаем блокировку. Все горутины заблокировались == deadlock.
	v := <-ch

	fmt.Println("Полученное из канала значение:", v)
}

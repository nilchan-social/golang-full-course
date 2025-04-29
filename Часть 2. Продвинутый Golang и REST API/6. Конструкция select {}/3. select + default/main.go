package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	// Через 500ms запишем в intCh значение
	go func() {
		time.Sleep(500 * time.Millisecond)
		intCh <- 1
	}()

	// Через 250ms запишем в strCh значение
	go func() {
		time.Sleep(250 * time.Millisecond)
		strCh <- "hi"
	}()

	// Засыпаем на 100ms в main горутине
	time.Sleep(100 * time.Millisecond)

	// На момент достижения конструкции select у нас ещё ни одни канал не будет доступен для чтения.
	// Отработает секция default
	select {
	case number := <-intCh:
		fmt.Println("Получено число:", number)
	case str := <-strCh:
		fmt.Println("Получена строка:", str)
	default:
		fmt.Println("На момент наступления секции select не было доступно ни одного канала для чтения. Отработала секция default.")
	}
}

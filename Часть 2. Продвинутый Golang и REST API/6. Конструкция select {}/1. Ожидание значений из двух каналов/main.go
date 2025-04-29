package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	// Через 200ms запишем в intCh значение
	go func() {
		time.Sleep(200 * time.Millisecond)
		intCh <- 1
	}()

	// Через 100ms запишем в strCh значение
	go func() {
		time.Sleep(100 * time.Millisecond)
		strCh <- "hi"
	}()

	// Блокируемя и ждём, когда в какой-то из каналов, intCh или strCh, придёт какое-то значение
	select {
	case number := <-intCh:
		fmt.Println("Получено число:", number)
	case str := <-strCh:
		fmt.Println("Получена строка:", str)
	}
}

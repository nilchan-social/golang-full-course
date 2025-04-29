package main

import (
	"fmt"
	"time"
)

// Обычная функция
func foo() {
	for {
		fmt.Println("Foo")
		time.Sleep(500 * time.Millisecond)
	}
}

// main горутина
func main() {
	// Запускаем обычную функцию в горутине
	go foo()

	// Запускаем анонимную функцию в горутине
	go func() {
		for {
			fmt.Println("Anon")
			time.Sleep(750 * time.Millisecond)
		}
	}()

	// Засыпаем на 5 секунд в main горутине, чтобы дать "Foo" и "Anon" горутинам поработать
	time.Sleep(5 * time.Second)
}

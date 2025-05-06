package main

import "fmt"

func main() {
	fmt.Println("До вызова функции hello()")
	hello()
	fmt.Println("После вызова функции hello()")
}

func hello() {
	fmt.Println("Я функция hello, я начинаюсь!")
	fmt.Println("Я фукнция hello, я продолжаюсь!")
	fmt.Println("Я функция hello, я заканчиваюсь!")
}

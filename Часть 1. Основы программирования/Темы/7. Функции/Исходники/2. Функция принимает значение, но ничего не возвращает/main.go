package main

import "fmt"

func main() {
	fmt.Println("До вызова функции square(x int)")
	square(5)
	fmt.Println("После вызова функции square(x int)")
}

func square(x int) {
	fmt.Println("Принято значение x:", x)
	fmt.Println("x в квадрате:", x*x)
}

package main

import "fmt"

func main() {
	score := 20

	if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться.")
	}
}

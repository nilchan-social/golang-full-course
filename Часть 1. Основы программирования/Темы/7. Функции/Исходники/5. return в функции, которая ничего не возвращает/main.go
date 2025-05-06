package main

import "fmt"

func main() {
	greeting("")
	greeting("Иван")
}

func greeting(name string) {
	if name == "" {
		fmt.Println("Вы передали пустое имя.")
		return
	}

	fmt.Println("Привет, уважаемый", name)
}

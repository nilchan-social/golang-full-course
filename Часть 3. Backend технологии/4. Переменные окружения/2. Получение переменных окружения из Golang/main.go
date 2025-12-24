package main

import (
	"fmt"
	"os"
)

func main() {
	some_var := os.Getenv("some_var")

	if some_var == "" {
		fmt.Println("Переменная окружения 'some_var' не задана.")
	} else {
		fmt.Println("Значение переменной окружения 'some_var':", some_var)
	}
}

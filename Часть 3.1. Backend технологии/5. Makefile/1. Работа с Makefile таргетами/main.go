package main

import (
	"fmt"
	"os"
)

func getEnvVar(key string) {
	v := os.Getenv(key)

	if v == "" {
		fmt.Printf("Переменная окружения '%s' не задана.\n", key)
	} else {
		fmt.Printf("Значение переменной окружения '%s': %s\n", key, v)
	}
}

func main() {
	fmt.Println("Запуск приложения!")
	getEnvVar("CONN_STRING")
	getEnvVar("SOME_ENV_VAR")
	fmt.Println("Завершение приложения.")
}

package main

import "fmt"

func main() {
	// 🔋
	// 🛢
	// 🐤

	fmt.Println("Начинаю генерировать уровень!")
	fmt.Println("")

	for i := 1; i <= 5; i++ {
		fmt.Println("Труба номер", i)
		fmt.Println("----")

		if i%2 == 0 {
			fmt.Println("🛢🛢")
		} else {
			fmt.Println("🔋🔋")
		}

		fmt.Println("----")
		fmt.Println("")
	}

	fmt.Println("Генерация уровня окончена!")
}

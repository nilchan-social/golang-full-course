package main

import "fmt"

func main() {
	// 🔋
	// 🐤

	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println("---------------------")
		fmt.Println("Вы подлетаете к трубе!", i)
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("Вы пролетаете через трубу!", i)
		fmt.Println("🔋🐤🔋")
		fmt.Println("")

		fmt.Println("Вы пролетели через трубу!", i)
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++

		fmt.Println("Счёт:", score)
		fmt.Println("")
	}
}

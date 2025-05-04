package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 🔋
	// 🐤
	// ❌

	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	// Бесконечный цикл (аналог while true в других языках)
	for {
		fmt.Println("---------------------")
		fmt.Println("Я подлетаю к трубе!")
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("Я пролетаю через трубу!")
		fmt.Println("🔋🐤🔋")
		fmt.Println("")

		if rand.Intn(8) == 1 {
			fmt.Println("Я врезался в трубу :(")
			fmt.Println("🔋❌🔋")

			// прерывание цикла
			break
		}

		fmt.Println("Я пролетел через трубу!")
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++

		fmt.Println("Счёт:", score)
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}
}

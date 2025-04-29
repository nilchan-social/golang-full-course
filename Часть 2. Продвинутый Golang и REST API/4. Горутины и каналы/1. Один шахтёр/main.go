package main

import (
	"fmt"
	"time"
)

// Фукнция, описывающая поход в шахту
// n -> номер похода в шахту
func mine(n int) int {
	// Имитируем поход в шахту на одну секунду
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	// Возвращаем добытый уголь
	return 10
}

// main горутина
func main() {
	coal := 0

	// Засекаем время
	initTime := time.Now()

	// Последовательно, раз за разом, ходим 3 раза в шахту
	coal += mine(1)
	coal += mine(2)
	coal += mine(3)

	// Итоговые значения угля и времени выполнения
	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}

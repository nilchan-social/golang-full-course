package main

import (
	"fmt"
	"time"
)

// Фукнция, описывающая поход в шахту
// transferPoint -> канал для передачи целочисленных значений между горутинами
// n             -> номер запущенной горутины
func mine(transferPoint chan int, n int) {
	// Имитируем поход в шахту на одну секунду
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	// Идём в пункт передачи угля
	// Если канал НЕ буферизированный, то будм тут блокироваться и ждать, пока кто-то из другой горутины примет передаваемое значение
	transferPoint <- 10

	// Сигнализируем о том, что значение было фактически передано через канал между двуми разными горутинами
	fmt.Println("Поход в шахту номер", n, "уголь успешно передан!")
}

// main горутина
func main() {
	coal := 0
	transferPoint := make(chan int)

	// Засекаем время
	initTime := time.Now()

	// Запускаем трёх шахтёров (запускаем 3 горутины)
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	// 3 раза приходим в пункт передачи угля за углём (3 раза читаем из канала)
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	// Итоговые значения угля и времени выполнения
	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}

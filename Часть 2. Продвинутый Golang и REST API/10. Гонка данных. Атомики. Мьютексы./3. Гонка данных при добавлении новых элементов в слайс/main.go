package main

import (
	"fmt"
	"sync"
)

// Глобалная переменная. Разделяемый между несколькими горутинами ресурс.
var slice []int

func addToSlice(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// Нерегулируемый конкурентный доступ к разделяемому ресурсу
		// Как следствие -- гонка данных
		slice = append(slice, i)
	}
}

func main() {
	// При помощи WaitGroup'ы дожидаемся окончательного выполнения 10-ти запущенных горутин
	wg := &sync.WaitGroup{}
	wg.Add(10)

	// Запускаем 10 горутин
	// Каждая 1.000 раз добавляет значение в слайс
	// Ожидается итоговый размер слайса == 10.000
	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)

	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)
	go addToSlice(wg)

	wg.Wait()

	// Из-за гонки данных мы потеряли много append'ов в слайс
	fmt.Println("Итоговый размер слайса:", len(slice))
}

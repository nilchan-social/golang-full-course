package main

import (
	"fmt"
	"sync"
)

// Глобалная переменная. Разделяемый между несколькими горутинами ресурс.
var number int = 0

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// Нерегулируемый конкурентный доступ к разделяемому ресурсу
		// Как следствие -- гонка данных
		number++
	}
}

func main() {
	// При помощи WaitGroup'ы дожидаемся окончательного выполнения 10-ти запущенных горутин
	wg := &sync.WaitGroup{}
	wg.Add(10)

	// Запускаем 10 горутин
	// Каждая 1.000 раз прибавляет 1 к number
	// Ожидается итоговое значение number == 10.000
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	// Из-за гонки данных мы потеряли много увеличений переменной number
	fmt.Println("Итоговое значение переменной number:", number)
}

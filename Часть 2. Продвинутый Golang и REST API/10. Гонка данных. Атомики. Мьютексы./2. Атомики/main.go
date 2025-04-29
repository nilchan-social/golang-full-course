package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Глобалная переменная. Разделяемый между несколькими горутинами ресурс.
var number atomic.Int64

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// Регулируем конкуретный доступ к целочисленной переменной благодаря атомикам
		// Гонка данных отсутствует
		number.Add(1)
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

	// Гонки данных нет. number == 10.000
	fmt.Println("Итоговое значение переменной number:", number.Load())
}

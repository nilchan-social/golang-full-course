package main

import (
	"fmt"
	"sync"
)

// Глобалная переменная. Разделяемый между несколькими горутинами ресурс.
var slice []int

// Мьютекст, позволяющий нам корректно регулировать конкуретный доступ к разделяемому ресурсу
var mtx sync.Mutex

func addToSlice(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// Мьютекс -- своего рода "эстафетная палочка"
		// Благодаря мьютексту, в критической секции горутины выстраиваются в очередь
		// Как следствие -- отсутствие гонки данных
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
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

	// Гонки данных нет, все 10.000 элементов будут успешно добавлены в слайс
	fmt.Println("Итоговый размер слайса:", len(slice))
}

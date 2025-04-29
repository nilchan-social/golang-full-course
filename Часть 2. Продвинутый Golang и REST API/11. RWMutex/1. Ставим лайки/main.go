package main

import (
	"fmt"
	"sync"
	"time"
)

// Глобалная переменная. Разделяемый между несколькими горутинами ресурс.
var likes int = 0

// Мьютекст, позволяющий нам корректно регулировать конкуретный доступ к разделяемому ресурсу
// Благодаря возможности отдельно брать блокировки на запись и отдельно брать блокировки на чтение
// Имеем значительный прирост в производительности приложения, при условии ощутимой нагрузки на чтение
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		// Регулируем конкуретную запись в переменную likes
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		// Регулируем конкуретное чтение из переменной likes
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()

	fmt.Println("Время выполнения:", time.Since(initTime))
}

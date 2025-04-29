package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция, описывающая почтальона, который 3 раза разносит указанную газету
func postman(wg *sync.WaitGroup, text string) {
	// В конце функции postman мы явно укажем WaitGroup'е, что функция postman завершилась
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнёс газету", text, "в", i, "раз")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Создаём WaitGroup
	wg := &sync.WaitGroup{}

	// Указываем WaitGroup'е, что "сейчас я запущу одну горутину"
	wg.Add(1)
	// Запускаем горутину
	go postman(wg, "Новости")

	// Указываем WaitGroup'е, что "сейчас я запущу одну горутину"
	wg.Add(1)
	// Запускаем горутину
	go postman(wg, "Игровой журнал")

	// Указываем WaitGroup'е, что "сейчас я запущу одну горутину"
	wg.Add(1)
	// Запускаем горутину
	go postman(wg, "Автомобильная хроника")

	// Дожидаемся, пока все запущенные горутины завершатся
	wg.Wait()

	fmt.Println("main завершился!")
}

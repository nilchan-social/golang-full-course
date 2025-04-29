package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция, описывающая работу одного отдельного почтальона
func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()

	for {
		// Этот почтально завершит своё выполнение, только после того
		// Как увидит сигнал о завершении рабочего дня через ctx
		// И доработает свою последнюю рабочую итерацию
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер:", n, "Мой рабочий день закончен!")
			return
		default:
			fmt.Println("Я почтальон номер:", n, "Взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон номер:", n, "Донёс письмо до почты:", mail)

			transferPoint <- mail

			fmt.Println("Я почтальон номер:", n, "Передал письмо:", mail)
		}
	}
}

// PostmanPool -- функция, запускающая postmanCount почтальонов,
// Позволяющая потребителям PostmanPool функции получать письма от запущенных почтальонов
// И контролирующая закрытие канала общения
func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

// Заготовка конкретных писем для конкретных почтальонов
// Если для какого-то почтальона нет заготовленного письма -- этот почтальон разносит лотерею
func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Семейный привет",
		2: "Приглашение от друга",
		3: "Информация из автосервиса",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}

	return mail
}

package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция, описывающая работу одного отдельного шахтёра
func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		// Этот шахтёр завершит своё выполнение, только после того
		// Как увидит сигнал о завершении рабочего дня через ctx
		// И доработает свою последнюю рабочую итерацию
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return
		default:
			fmt.Println("Я шахтёр номер:", n, "Начал добывать уголь!")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтёр номер:", n, "Добыл уголь:", power)

			transferPoint <- power
			fmt.Println("Я шахтёр номер:", n, "Передал уголь:", power)
		}
	}
}

/*
func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		// Этот шахтёр завершит своё выполнение СРАЗУ после сигнала о завершении рабочего дня
		fmt.Println("Я шахтёр номер:", n, "Начал добывать уголь!")

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтёр номер:", n, "Добыл уголь:", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтёр номер:", n, "Передал уголь:", power)
		}

	}
}
*/

// MinerPool -- функция, запускающая minerCount шахтёров,
// Позволяющая потребителям MinerPool функции получать уголь от запущенных шахтёров
// И контролирующая закрытие канала общения
func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}

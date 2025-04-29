package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Хранилище угля
	var coal atomic.Int64

	// Хранилище писем
	mtx := sync.Mutex{}
	var mails []string

	// Создаём отдельные контексты для шахтёров и почтальонов
	// Для возможности отдельно завершать шахтёров и отдельно завершать почтальонов
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	// Через 3 секунды мы завершим рабочий день шахтёров
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->>> Рабочий день шахтёров окончен!")
		minerCancel()
	}()

	// Через 6 секунд мы завершим рабочий день почтальонов
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->>> Рабочий день почтальоном окончен!")
		postmanCancel()
	}()

	// Запускаем 2-х шахтёров, получаем пункт передачи угля
	coalTransferPoint := miner.MinerPool(minerContext, 2)
	// Запускаем 2-х почтальонов, получаем пункт передачи писем
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	// Засекаем время выполнения всей работы
	initTime := time.Now()

	wg := &sync.WaitGroup{}

	// В отдельной горутине вычитываем входящий уголь
	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
			time.Sleep(1 * time.Second)
		}
	}()

	// В отдельной горутине вычитываем входящие письма
	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()

			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()

	// Далее выводим результирующие значения

	fmt.Println("Суммарно добытый уголь:", coal.Load())

	mtx.Lock()
	fmt.Println("Суммарное количество полученных писем:", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время:", time.Since(initTime))
}

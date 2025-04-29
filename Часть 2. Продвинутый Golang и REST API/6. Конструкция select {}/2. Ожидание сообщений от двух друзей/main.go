package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	// Канал передачи сообщений от первого друга
	messageChan1 := make(chan Message)
	// Канал передачи сообщений от второго друга
	messageChan2 := make(chan Message)

	// В одной горутине запускаем первого "Друга", который нам раз в 3 секунды посылает сообщение
	go func() {
		for {
			messageChan1 <- Message{
				Author: "Друг 1",
				Text:   "Привет",
			}

			time.Sleep(3 * time.Second)
		}
	}()

	// Во второй горутине запускаем второго "Друга", который нам раз в 250 миллисекунд посылает сообщение
	go func() {
		for {
			messageChan2 <- Message{
				Author: "Друг 2",
				Text:   "Как дела?",
			}

			time.Sleep(250 * time.Millisecond)
		}
	}()

	// Бесконечное количество раз
	for {
		// Ожидаем сообщение от одного из каналов
		select {
		case msg1 := <-messageChan1:
			fmt.Println("Я получил сообщение от:", msg1.Author, "Текст сообщения:", msg1.Text)
		case msg2 := <-messageChan2:
			fmt.Println("Я получил сообщение от:", msg2.Author, "Текст сообщения:", msg2.Text)
		}
	}
}

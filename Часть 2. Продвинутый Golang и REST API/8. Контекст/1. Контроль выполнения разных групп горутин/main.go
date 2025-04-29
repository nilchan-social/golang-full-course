package main

import (
	"context"
	"fmt"
	"time"
)

func Foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo завершилась :(", n)
			return
		default:
			fmt.Println("Foo продолжает своё выполнение!", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func Boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo завершилась :(", n)
			return
		default:
			fmt.Println("Boo продолжает своё выполнение!", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// создаём родительский контекст
	parentContext, parentCancel := context.WithCancel(context.Background())
	// наследуем от родительского контекста дочерний контекст
	childContext, childCancel := context.WithCancel(parentContext)

	// запускаем 3 горутины с родительским контекстом
	go Foo(parentContext, 1)
	go Foo(parentContext, 2)
	go Foo(parentContext, 3)

	// запускаем 3 горутины с дочерним контекстом
	go Boo(childContext, 1)
	go Boo(childContext, 2)
	go Boo(childContext, 3)

	// через 2 секунды отменяем дочерний контекст
	// тем самым завершая все горутины, которые зависят от дочернего контекста
	time.Sleep(2 * time.Second)
	childCancel()

	// через ещё 2 секунды завершаем родительский конткст
	// завершая все горутины, которые зависят от родительского контекста
	// а так же автоматически завершается дочерний контекст, если он ещё не был отменён
	time.Sleep(2 * time.Second)
	parentCancel()

	// ещё одну секунду спим в main горутине, чтобы продемонстрировать
	// что Foo и Boo горутины завершились именно при помощи контекста
	// а просто вместе со всей программой
	time.Sleep(1 * time.Second)
	fmt.Println("main завершился!")
}

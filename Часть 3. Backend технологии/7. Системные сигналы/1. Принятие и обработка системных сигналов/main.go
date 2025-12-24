package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("PID:", os.Getpid())

	sigintChan := make(chan os.Signal, 1)
	sigtermChan := make(chan os.Signal, 1)
	sigkillChan := make(chan os.Signal, 1)

	signal.Notify(sigintChan, syscall.SIGINT)
	signal.Notify(sigtermChan, syscall.SIGTERM)
	signal.Notify(sigkillChan, syscall.SIGKILL)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		i := 0
		for sig := range sigintChan {
			i++
			fmt.Println("SIGNAL:", sig, i)

			if i == 3 {
				fmt.Println("SIGINT received 3 times, stop receiving SIGINT...")
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		i := 0
		for sig := range sigtermChan {
			i++
			fmt.Println("SIGNAL:", sig, i)

			if i == 3 {
				fmt.Println("SIGTERM received 3 times, stop receiving SIGTERM...")
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		i := 0
		for sig := range sigkillChan {
			i++
			fmt.Println("SIGNAL:", sig, i)

			if i == 3 {
				fmt.Println("SIGKILL received 3 times, stop receiving SIGKILL...")
				return
			}
		}
	}()

	fmt.Println("Waiting for SIGINT/SIGTERM/SIGKILL system signals...")
	wg.Wait()
	fmt.Println("Each signal (SIGINT/SIGTERM) received 3 times, program stopped.")
}

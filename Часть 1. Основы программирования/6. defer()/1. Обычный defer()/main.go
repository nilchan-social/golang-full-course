package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer Println")
	}()

	fmt.Println("hello")
}

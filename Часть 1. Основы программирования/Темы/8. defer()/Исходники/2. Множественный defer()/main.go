package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer1")
	}()

	defer func() {
		fmt.Println("Defer2")
	}()

	defer func() {
		fmt.Println("Defer3")
	}()

	fmt.Println("hello")
}

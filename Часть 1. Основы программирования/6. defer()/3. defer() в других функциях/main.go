package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("main Defer1")
	}()
	defer func() {
		fmt.Println("main Defer2")
	}()
	defer func() {
		fmt.Println("main Defer3")
	}()

	fmt.Println("до вызова foo")
	foo()
	fmt.Println("после вызова foo")
}

func foo() {
	defer func() {
		fmt.Println("foo Defer1")
	}()

	defer func() {
		fmt.Println("foo Defer2")
	}()

	defer func() {
		fmt.Println("foo Defer3")
	}()

	fmt.Println("foo")
}

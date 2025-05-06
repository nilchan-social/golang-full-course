package main

import "fmt"

func main() {
	number := 15

	var notNilPtr *int = &number
	var nilPtr *int

	fmt.Print("Not nil pointer: ")
	foo(notNilPtr)

	fmt.Println("")

	fmt.Print("nil pointer: ")
	foo(nilPtr)
}

func foo(intPtr *int) {
	fmt.Println("Указатель:", intPtr)

	if intPtr != nil {
		fmt.Println("Значение переменной, на которую указывает указатель:", *intPtr)
	} else {
		fmt.Println("Указатель является nil-указателем, а значит я не могу его безопасно разыменовать")
	}
}

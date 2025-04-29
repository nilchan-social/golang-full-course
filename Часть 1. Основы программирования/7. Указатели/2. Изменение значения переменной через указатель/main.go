package main

import "fmt"

func main() {
	number := 5
	numberPtr := &number

	fmt.Println("number до вызова foo:", number)
	foo(numberPtr)
	fmt.Println("number после вызова foo:", number)

	text := "hello"
	textPtr := &text

	fmt.Println("text до вызова foo:", text)
	boo(textPtr)
	fmt.Println("text после вызова foo:", text)
}

func foo(intPtr *int) {
	*intPtr = 10
}

func boo(stringPtr *string) {
	*stringPtr = "world"
}

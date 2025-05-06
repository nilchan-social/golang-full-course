package main

import "fmt"

func main() {
	number := 5

	ptr := &number

	foo(ptr)
}

func foo(intPtr *int) {
	fmt.Println("Указатель:", intPtr)
	fmt.Println("Значение переменной, на которую указатель указывает:", *intPtr)
}

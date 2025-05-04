package main

import "fmt"

// 10
var number int = 10

func main() {
	// 15
	number += 5

	foo()
	boo()

	fmt.Println("number:", number)
}

func foo() {
	// 30
	number *= 2
}

func boo() {
	// 3
	number /= 10
}

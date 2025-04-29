package main

import "fmt"

func main() {
	number := 5
	text := "hello"

	fmt.Println("number до:", number)
	fmt.Println("text до:", text)

	foo(number, text)

	fmt.Println("number после:", number)
	fmt.Println("text после:", text)

}

func foo(n int, t string) {
	n = 10
	t = "world"
}

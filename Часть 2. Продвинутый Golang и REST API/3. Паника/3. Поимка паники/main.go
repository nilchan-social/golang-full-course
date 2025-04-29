package main

import "fmt"

func panicSlice() {
	slice := []int{11, 22, 33}

	i := slice[3]

	fmt.Println("i:", i)
}

func panicDivision() {
	number := 0

	i := 10 / number

	fmt.Println("i:", i)
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была отловлена паника:", p)
		}
	}()

	// panicSlice()
	// panicDivision()

	fmt.Println("Конец выполнения программы!")
}

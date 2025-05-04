package main

import "fmt"

func main() {
	number := 15
	text := "Hello"
	drob := 11.25
	boolean := true

	// сложение
	number = number + 10
	number += 15
	number++

	// умножение
	number = number * 100
	number *= 5

	// деление
	number = number / 10
	number /= 5

	// вычитание
	drob = drob - 1.05
	drob -= 4.2
	drob--

	// конкатенация
	text = text + "World"
	text += "!"

	// вывод в консоль
	fmt.Println("number:", number)
	fmt.Println("text:", text)
	fmt.Println("drob:", drob)
	fmt.Println("boolean:", boolean)
}

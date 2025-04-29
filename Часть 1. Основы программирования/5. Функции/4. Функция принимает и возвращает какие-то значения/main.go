package main

import "fmt"

func main() {
	number1 := 25
	number2 := 17

	resultNumber := sumInt(number1, number2)

	fmt.Println("number1 + number2 ==", resultNumber)

	text1 := "hello"
	text2 := "world"

	resultText := sumText(text1, text2)

	fmt.Println("text1 + text2 ==", resultText)
}

func sumInt(a int, b int) int {
	s := a + b

	return s
}

func sumText(s1 string, s2 string) string {
	s := s1 + s2

	return s
}

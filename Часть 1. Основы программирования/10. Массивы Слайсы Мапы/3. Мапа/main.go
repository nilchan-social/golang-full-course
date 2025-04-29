package main

import "fmt"

func main() {
	foo1()
	fmt.Println("-----------")
	foo2()
}

func foo1() {
	weather := map[int]int{
		10: +3,
		11: -4,
		12: +1,
	}

	weather[12] = -10
	weather[13] = +5

	for key, value := range weather {
		fmt.Println("key:", key, ", value:", value)
	}

	w11 := weather[11]
	fmt.Println("погода 11-го числа:", w11)
}

func foo2() {
	weather := make(map[int]int)

	weather[10] = -2
	weather[11] = -1
	weather[12] = 0

	w12, ok12 := weather[12]
	w13, ok13 := weather[13]

	fmt.Println("погода 12-го числа:", w12, ", является ли это записанным ранее значением, а не значением по-умолчанию:", ok12)
	fmt.Println("погода 13-го числа:", w13, ", является ли это записанным ранее значением, а не значением по-умолчанию:", ok13)
}

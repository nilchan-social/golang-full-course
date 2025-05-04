package main

import "fmt"

func main() {
	sunny := true
	weekend := true

	// Я пойду гулять, только если будет солнечная погода и будет выходной,
	// иначе я никуда не пойду

	if sunny && weekend {
		fmt.Println("Я иду гулять!")
	} else {
		fmt.Println("Я НЕ пойду гулять (")
	}
}

package main

import "fmt"

func main() {
	computerClub := true
	icecream := true

	// Я пойду гулять, если там либо будет компьютерный клуб
	// либо если там будет мороженное.
	// Если там не будет ни того ни другого, я никуда не пойду

	if computerClub || icecream {
		fmt.Println("Я иду гулять!")
	} else {
		fmt.Println("Я НЕ пойду гулять (")
	}
}

package main

import (
	"fmt"
	"study/greeting"
	"study/greeting/foo"
	"study/user"

	"github.com/k0kubun/pp"
)

func main() {
	greeting.SayHello()
	i := greeting.GiveMeInt()
	fmt.Println("i:", i)

	rand := foo.RandomFunc()
	fmt.Println("rand:", rand)

	u := user.User{
		Name: "Григорий",
		// age не видно, потому что age private
	}

	u.ChangeAge(50)
	// changeName не видно, потому что changeName private
	// user.changeName

	pp.Println("user:", u)
}

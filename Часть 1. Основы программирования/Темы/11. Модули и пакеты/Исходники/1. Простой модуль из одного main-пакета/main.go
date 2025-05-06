package main

import "github.com/k0kubun/pp"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Василий",
		Age:  67,
	}

	pp.Println(user)
}

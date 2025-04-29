package main

import "fmt"

type User struct {
	Age int
}

func (u User) ChangeAge1(newAge int) {
	u.Age = newAge
}

func (u *User) ChangeAge2(newAge int) {
	u.Age = newAge
}

func main() {
	user := User{
		Age: 30,
	}

	fmt.Println("user изначально:", user)
	user.ChangeAge1(35)
	fmt.Println("user после ChangeAge1:", user)
	user.ChangeAge2(35)
	fmt.Println("user после ChangeAge2:", user)
}

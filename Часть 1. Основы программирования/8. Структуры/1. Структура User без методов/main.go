package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func main() {
	user1 := User{
		Name:        "Иван",
		Age:         50,
		PhoneNumber: "+7 (911) 911 91-91",
		IsClose:     true,
		Rating:      5.5,
	}

	user2 := User{
		Name:        "Пётр",
		Age:         60,
		PhoneNumber: "+7 (922) 922 92-92",
		// IsClose не указано => значение по умолчанию (false)
		// Rating не указано => значение по умолчанию (0.0)
	}

	fmt.Println("user1:", user1)
	fmt.Println("user1 Name:", user1.Name)

	fmt.Println("user2:", user2)
	fmt.Println("user2 Rating:", user2.Rating)
}

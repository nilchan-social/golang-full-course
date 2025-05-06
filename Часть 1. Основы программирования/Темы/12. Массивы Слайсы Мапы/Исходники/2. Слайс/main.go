package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	foo1()
	fmt.Println("<<<<<<<>>>>>>>")
	foo2()
	fmt.Println("<<<<<<<>>>>>>>")
	foo3()
}

func foo1() {
	slice1 := []User{
		User{
			Name:    "Иван",
			Rating:  5.5,
			Premium: true,
		},
		User{
			Name:    "Сергей",
			Rating:  4.5,
			Premium: false,
		},
		User{
			Name:    "Данил",
			Rating:  7.5,
			Premium: true,
		},
	}

	pp.Println("slice1 до:", slice1)
	fmt.Println("---------------")

	// можем динамически добавлять элементы
	slice1 = append(
		slice1, User{
			Name:    "Олег",
			Rating:  8.2,
			Premium: false,
		},
	)
	slice1 = append(
		slice1, User{
			Name:    "Пётр",
			Rating:  3.1,
			Premium: true,
		},
	)

	pp.Println("slice1 после:", slice1)
	fmt.Println("---------------")
}

func foo2() {
	slice2 := make([]User, 5)
	pp.Println("slice2:", slice2)
}

func foo3() {
	// можем заранее предвыделить память
	// если знаем кол-во элементов
	// make([]User, 0, 3)
	slice3 := make([]User, 0)

	slice3 = append(slice3, User{
		Name:    "Иван",
		Rating:  5.5,
		Premium: true,
	})

	slice3 = append(slice3,
		User{
			Name:    "Сергей",
			Rating:  4.5,
			Premium: false,
		},
		User{
			Name:    "Данил",
			Rating:  7.5,
			Premium: true,
		},
	)

	pp.Println("slice3 до:", slice3)
	fmt.Println("---------------")

	// каждому у кого Premium аккаун добавляет 1.0 соц рейтинга
	for index, user := range slice3 {
		if user.Premium {
			slice3[index].Rating += 1.0
		}
	}

	pp.Println("slice3 после:", slice3)
	fmt.Println("---------------")
}

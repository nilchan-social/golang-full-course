package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	// Проверяем, хватает ли на балансе пользователя денег для совершения покупки
	if user.Ballance-usd < 0 {
		// Если не хватает, то возвращаем ошибку оплаты с сообщением "Недостаточно средств"
		return errors.New("Недостаточно средств!")
	}

	// Если хватает, то проводим оплату, списываем деньги
	user.Ballance -= usd

	// Ошибок не было => возвращаем nil (nil error == отсутствие ошибки)
	return nil
}

func main() {
	ballance := 100
	payment := 50

	/*
		- А если?

		ballance := 30
		payment := 50
	*/

	user := User{
		Name:     "Олег",
		Ballance: ballance,
	}

	pp.Println("User до проведения оплаты:", user)
	err := Pay(&user, payment)
	pp.Println("User после проведения оплаты:", user)

	if err != nil {
		fmt.Println("Оплаты не было! Причина:", err)
	} else {
		fmt.Println("Была произведена оплата! Ошибок не было.")
	}
}

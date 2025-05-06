package main

import "fmt"

type User struct {
	// Имя пользователя
	// Правило: имя должно быть не пустое
	Name string

	// Возраст пользователя
	// Правило: возраст должен быть больше нуля и меньше 150
	Age int

	// Номер телефона пользователя
	// Правило: номер телефона должен быть не пустым
	PhoneNumber string

	// Закрыт ли профиль пользователя?
	IsClose bool

	// Социальный рейтинг пользователя
	// Правило: рейтинг должен быть больше либо равен нулю, но меньше либо равен десяти
	Rating float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	fmt.Println("Валидирую имя!")
	if name == "" {
		fmt.Println("Имя не прошло валидацию :(")
		return User{}
	}
	fmt.Println("Имя прошло валидацию!")

	fmt.Println("Валидирую возраст!")
	if age <= 0 || age >= 150 {
		fmt.Println("Возраст не прошёл валидацию :(")
		return User{}
	}
	fmt.Println("Возраст прошёл валидацию!")

	fmt.Println("Валидирую номер телефона!")
	if phoneNumber == "" {
		fmt.Println("Номер телефона не прошёл валидацию :(")
		return User{}
	}
	fmt.Println("Номер телефона прошёл валидацию!")

	fmt.Println("Валидирую рейтинг!")
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Рейтинг не прошёл валидацию :(")
		return User{}
	}
	fmt.Println("Рейтинг прошёл валидацию!")

	fmt.Println("Валидация в конструкторе успешно завершена!")
	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	// закрыть аккаунт
	u.IsClose = true
}

func (u *User) OpenAccount() {
	// открыть аккаунт
	u.IsClose = false
}

func (u *User) RatingUp(ratingDiff float64) {
	if u.Rating+ratingDiff <= 10.0 {
		u.Rating += ratingDiff
	}
}

func (u *User) RatingDown(ratingDiff float64) {
	if u.Rating-ratingDiff >= 0.0 {
		u.Rating -= ratingDiff
	}
}

func main() {
	user1 := NewUser(
		"Иван",
		50,
		"+7 (911) 911 91-91",
		true,
		5.5,
	)

	user2 := NewUser(
		"Пётр",
		60,
		"+7 (922) 922 92-92",
		false,
		500.5,
	)

	user1.RatingUp(1.5)
	user2.CloseAccount()

	fmt.Println("user1:", user1)
	fmt.Println("user2:", user2)
}

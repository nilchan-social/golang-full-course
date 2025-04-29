package main

import "fmt"

// Интерфейс Auto
// Каждая структура, у которой есть методы:
// - StepOnGas()
// - StepOnBrake()
// Автоматически удовлетворяет интерфейсу Auto и может быть использована на его месте
type Auto interface {
	StepOnGas()
	StepOnBrake()
}

// Структура BMW
// Имеет:
// - Метод StepOnGas()
// - Метод StepOnBrake()
// - Метод BipBip()
// Вывод:
// - Наличие методов StepOnGas() + StepOnBrake() говорит о том, что BMW подходит под интерфейс Auto
// - Наличие метода BipBip() не мешает структуре BMW удовлетворять интерфейсу Auto
type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ! 550 лошадинных сил! Жмём на газ!")
}

func (b BMW) StepOnBrake() {
	fmt.Println("Я БМВ! Тормоз в пол! Хорошее торможение!")
}

func (b BMW) BipBip() {
	fmt.Println("Я БМВ! Я нажимаю клаксон!")
}

// Структура Zhiga
// Имеет:
// - Метод StepOnGas()
// - Метод StepOnBrake()
// Вывод:
// - Наличие методов StepOnGas() + StepOnBrake() говорит о том, что Zhiga подходит под интерфейс Auto
type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Я жига! 2107! Пробую не развалиться!")
}

func (z Zhiga) StepOnBrake() {
	fmt.Println("Я жига! Вот блин тормоза чуть не отвалились(")
}

// Структура Lambo
// Имеет:
// - Метод StepOnGas()
// Вывод:
// - Отсутствие метода StepOnBreak() не позволяет структуре Lambo удовлетворять интерфейсу Auto
type Lambo struct{}

func (l Lambo) StepOnGas() {
	fmt.Println("Я ламба! 1500 лс! Еду быстро!")
}

// Функция ride(auto Auto)
// Принимает:
// - interface Auto
// Вывод:
// - Функция ride(auto Auto) может принимать любую структуру, которая подходит под интерфейс Auto
func ride(auto Auto) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю на газ...")
	auto.StepOnGas()
	auto.StepOnBrake()
}

func main() {
	/*
		bmw := BMW{}
		zhiga := Zhiga{}
		lambo := Lambo{}

		ride(bmw)   // ok
		ride(zhiga) // ok
		ride(lambo) // ошибка: в структуре Lambo отсутсвует необходимый метод StepOnBrake(), Lambo не является Auto
	*/
}

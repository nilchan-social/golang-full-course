package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	// Проверили: есть ли ещё прочность у автомобиля, чтобы газануть?
	if c.Armor-10 <= 0 {
		// Если прочности нет, то мы не газуем, а возвращаем ошибку
		return 0, errors.New("Мы не стали газовать, чтобы не сломать машину!")
	}

	// Газуем, получаем случайное количество километров в час разгона, уменьшаем остаточную мощность автомобиля
	kmch := rand.Intn(150)
	c.Armor -= 10

	// Возвращаем новый разгон, ошибок нет
	return kmch, nil
}

func main() {
	car := Car{
		Armor: 25,
	}

	for {
		pp.Println("Car до нажатия на газ:", car)
		kmch, err := car.Gas()
		pp.Println("Car после нажатия на газ:", car)
		if err != nil {
			fmt.Println("Ошибка нажатия на газ:", err)
			break
		}

		fmt.Println("Получившийся разгон:", kmch)
		fmt.Println("")
	}
}

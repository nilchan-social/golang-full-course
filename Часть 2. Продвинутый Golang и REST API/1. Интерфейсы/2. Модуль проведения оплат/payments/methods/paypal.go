package methods

import (
	"fmt"
	"math/rand"
)

// Заглушка, иммитирующая проведение оплаты через PayPal
type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (p PayPal) Pay(usd int) int {
	// Иммитируем проведения оплаты через PayPal
	fmt.Println("Оплата через PayPal!")
	fmt.Println("Размер оплаты:", usd, "usd")

	// Генерируем и возвращаем рандомный ID оплаты
	return rand.Int()
}

func (p PayPal) Cancel(id int) {
	// Иммитируем отмену оплаты с переданным ID через PayPal
	fmt.Println("Отмена операции PayPal! ID:", id)
}

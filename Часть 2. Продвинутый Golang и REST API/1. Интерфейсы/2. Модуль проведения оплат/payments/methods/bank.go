package methods

import (
	"fmt"
	"math/rand"
)

// Заглушка, иммитирующая проведение оплаты через банк
type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	// Иммитируем проведения оплаты через банк
	fmt.Println("Оплата через банк!")
	fmt.Println("Размер оплаты:", usd, "долларов")

	// Генерируем и возвращаем рандомный ID оплаты
	return rand.Int()
}

func (b Bank) Cancel(id int) {
	// Иммитируем отмену оплаты с переданным ID через банк
	fmt.Println("Отмена банковской операции! ID:", id)
}

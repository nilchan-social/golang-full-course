package methods

import (
	"fmt"
	"math/rand"
)

// Заглушка, иммитирующая проведение оплаты через крипто-кошелёк
type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	// Иммитируем проведения оплаты через крипто-кошелёк
	fmt.Println("Оплата криптовалютой!")
	fmt.Println("Размер оплаты:", usd, "USDT")

	// Генерируем и возвращаем рандомный ID оплаты
	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	// Иммитируем отмену оплаты с переданным ID через крипто-кошелёк
	fmt.Println("Отмена крипто-операци! ID:", id)
}

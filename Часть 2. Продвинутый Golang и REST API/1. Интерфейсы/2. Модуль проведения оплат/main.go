package main

import (
	"interfaces/payments"
	"interfaces/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	// Выбираем метод оплаты
	// method := methods.NewBank()   // Оплата через банк
	// method := methods.NewPayPal() // Оплата через PayPal
	method := methods.NewCrypto() // Оплата через крипто-кошелёк

	// Создаём модуль проведения оплат
	paymentModule := payments.NewPaymentModule(method)

	// Проводим 3 оплаты
	paymentModule.Pay("Бургер", 5)
	phonePaymentID := paymentModule.Pay("Телефон", 500)
	paymentModule.Pay("Игра", 20)

	// По сохранённому ID оплаты за телефон, мы отменяем эту самую оплату
	paymentModule.Cancel(phonePaymentID)

	// Получаем информацию по всем проведённым оплатам
	allInfo := paymentModule.AllInfo()

	// Выводим эту информацию в консоль
	pp.Println("Информация по всем проведённым оплатам:", allInfo)
}

package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
	paymentsInfo  map[int]PaymentsInfo
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
		paymentsInfo:  make(map[int]PaymentsInfo),
	}
}

// Метод Pay()
// Принимает:
// - Описание проводимой оплаты
// - Сумму оплаты
// Возвращает:
// - ID проведённой операции
func (p *PaymentModule) Pay(description string, usd int) int {
	// Проводим оплату через выбранный метод проведения оплат
	// В ответ получаем ID проведённой оплаты
	id := p.paymentMethod.Pay(usd)

	// Собираем всю информацию об оплате в одну структуру
	info := PaymentsInfo{
		Description: description,
		USD:         usd,
		Cancelled:   false,
	}

	// Сохраняем собранную информацию об оплате
	// В качестве ключа: ID проведённой операции
	p.paymentsInfo[id] = info

	// Возвращаем ID операции
	return id
}

// Метод Cancel()
// Принимает:
// - ID операции
// Возвращает:
// - Ничего не возвращает
func (p *PaymentModule) Cancel(id int) {
	// Получаем информацию об оплате по переданному ID
	info, ok := p.paymentsInfo[id]
	// Проверяем, была ли вообще ранее проведена оплата с переданным ID
	if !ok {
		// Если оплата не была проведена, то просто завершаем выполнение текущего метода
		return
	}

	// Если оплата была проведена, то отменяем её через выбранный метод проведения оплат
	p.paymentMethod.Cancel(id)

	// Сохраняем информацию о том, то мы отменили оплату с определённым ID
	info.Cancelled = true
	p.paymentsInfo[id] = info
}

// Метод Info()
// Принимает:
// - ID операции
// Возвращает:
// - Информацию о проведённой операции
func (p *PaymentModule) Info(id int) PaymentsInfo {
	// Получаем информацию об оплате по переданному ID
	info, ok := p.paymentsInfo[id]
	// Проверяем, была ли вообще ранее проведена оплата с переданным ID
	if !ok {
		// Если оплата не была проведена, то просто возвращаем пустую информацию об оплате
		return PaymentsInfo{}
	}

	// Если оплата была проведена, то возвращаем найденную информацию об этой операции
	return info
}

// Метод AllInfo()
// Принимает:
// - Ничего не принимает
// Возвращает:
// - Информацию о всех проведённых операциях
func (p *PaymentModule) AllInfo() map[int]PaymentsInfo {
	// Создаём временную мапу tmp, чтобы защитить мапу p.paymentsInfo от внешнего воздействия потребителями метода AllInfo()
	tmp := make(map[int]PaymentsInfo, len(p.paymentsInfo))

	// Перекопирукм во временную мапу все значения из основной мапы
	for k, v := range p.paymentsInfo {
		tmp[k] = v
	}

	// Возвращаем временную мапу, в которой лежат скопированные из основной мапы данные
	return tmp
}

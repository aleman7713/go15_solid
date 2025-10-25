package notification

import (
	"fmt"
)

// Структура заказа
type SMSSender struct {
	ID int
	Customer string
	Products string
	Total float64
	Status string
}

func (v SMSSender) Send(type1 int) {
	switch type1 {
	case 1:
		fmt.Printf("SMSSender: Уведомление отправлено клиенту %s\n", v.Customer)
	case 2:
		fmt.Printf("SMSSender: Уведомление отправлено клиенту %s o заказе №%d\n", v.Customer, v.ID)
	case 3:
		fmt.Printf("SMSSender: Уведомление отправлено клиенту %s o заказе №%d на сумму %.2f\n", v.Customer, v.ID, v.Total)
	}
}
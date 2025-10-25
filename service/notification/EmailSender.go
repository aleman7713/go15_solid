package notification

import (
	"fmt"
)

// Структура заказа
type EmailSender struct {
	Customer string
	Products string
	Total float64
}

func (v EmailSender) Send(type1 int, id int) {
	switch type1 {
	case 1:
		fmt.Printf("EmailSender: Уведомление отправлено клиенту %s\n", v.Customer)
	case 2:
		fmt.Printf("EmailSender: Уведомление отправлено клиенту %s o заказе №%d\n", v.Customer, id)
	case 3:
		fmt.Printf("EmailSender: Уведомление отправлено клиенту %s o заказе №%d на сумму %.2f\n", v.Customer, id, v.Total)
	}
}
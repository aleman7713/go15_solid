package postgres

import (
	"fmt"
	db1 "go15_solid/repository/db"
	n "go15_solid/service/notification"
	"strings"
)

type OrderSystem struct {
	Db db1.RepositoryWriter
}

func NewOrderSystem(db db1.RepositoryWriter) *OrderSystem {
	return &OrderSystem{Db: db}
}

var id int = 0

// Работа с БД для заказов
func (v *OrderSystem) CreateOrder(customer string, products []string, total float64) error {
	id++
	status := "pending"

	// Создание заказа в БД
	sql := fmt.Sprintf("INSERT INTO orders (customer, products, total, status) VALUES ('%s', '%v', %.2f, '%s')",
		customer, products, total, status)
	err := v.Db.Exec(sql)

	if err != nil {
		return err
	}

	// Отправка уведомления
	var v2 n.Notifier = n.EmailSender{
		ID:       id,
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total:    total,
		Status:   status,
	}
	v2.Send(1)

	v2 = n.SMSSender{
		ID:       id,
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total:    total,
		Status:   status,
	}
	v2.Send(1)
	
	return nil
}

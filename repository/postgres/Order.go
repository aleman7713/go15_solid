package postgres;

import (
	"fmt"
	"strings"
	n "go15_solid/service/notification"
	m "go15_solid/repository/model"
	db1 "go15_solid/repository/db"
)

type OrderSystem struct {
	Db db1.Idb
}

func NewOrderSystem(db db1.Idb) *OrderSystem {
	return &OrderSystem{Db: db}
}

var id int = 0

// Работа с БД для заказов
func (v *OrderSystem) CreateOrder(customer string, products []string, total float64) error {
	id++
	status := "pending"

	// Создание заказа в БД
	sql := fmt.Sprintf("INSERT INTO orders (customer, products, total, status) VALUES ('%s', '%v', %.2f, '%s')",
		customer, products, total, status,)
	err := v.Db.Exec(sql)

	if err != nil {
		return err
	}

	// Отправка уведомления
	var v2 n.Notification = m.Order{
		ID: id,
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total: total,
		Status: status,
	}
	v2.Send(1)

	return nil		
}

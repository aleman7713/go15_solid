package postgres

import (
	db1 "go15_solid/repository/db"
	n "go15_solid/service/notification"
)

type OrderSystem struct {
	Db db1.RepositoryWriter
    Notifier n.Notifier
}

func NewOrderSystem(db db1.RepositoryWriter, notifier n.Notifier) *OrderSystem {
	return &OrderSystem{Db: db, Notifier: notifier}
}

var id int = 0

// Работа с БД для заказов
func (v *OrderSystem) CreateOrder(customer string, products []string, total float64) error {
	id++
	status := "pending"

	// Создание заказа в БД
	err := v.Db.ExecOrder(customer, products, total, status)

	if err != nil {
		return err
	}

	// Отправка уведомления
	v.Notifier.Send(1, id)

	return nil
}

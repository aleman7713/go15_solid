package main

import (
	"strings"
	db1 "go15_solid/repository/db"
	"go15_solid/repository/postgres"
    n "go15_solid/service/notification"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db db1.RepositoryWriter = &db1.DB{}

	err := db.Open("sqlite3", "orders.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer TEXT NOT NULL,
			products TEXT NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}


	// Создаём заказ с оповещением на email
	customer := "Иван"
	products := []string{"apple", "banana"}
	total := 10.5

	var v1 n.Notifier = n.EmailSender{
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total:    total,
	}

	system_email := postgres.NewOrderSystem(db, v1)

	err = system_email.CreateOrder(customer, products, total)
	if err != nil {
		log.Fatal(err)
	}

	// Создаём заказ с оповещением на SMS
	customer = "Алексей"
	products = []string{"juice", "salad"}
	total = 7.5

	var v2 = n.SMSSender{
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total:    total,
	}

	system_sms := postgres.NewOrderSystem(db, v2)

	err = system_sms.CreateOrder(customer, products, total)
	if err != nil {
		log.Fatal(err)
	}
}

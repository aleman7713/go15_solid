package main

import (
	//	"database/sql"
	db1 "go15_solid/repository/db"
	"go15_solid/repository/postgres"
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

	system := postgres.NewOrderSystem(db)

	err = system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}
}

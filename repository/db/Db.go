package db

import (
	"database/sql"
	"fmt"
)

type RepositoryWriter interface {
	Open(driverName string, dataSourceName string) error
	Close() error
	Exec(sql string) error
	ExecOrder(customer string, products []string, total float64, status string) error
}

type DB struct {
	Db_type string
	Db      *sql.DB
}

func (v *DB) Close() error {
	err := v.Db.Close()
	return err
}

func (v *DB) ExecOrder(customer string, products []string, total float64, status string) error {
	_, err := v.Db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		customer, fmt.Sprintf("%v", products), total, status,
	)
	return err
}

func (v *DB) Exec(sql string) error {
	_, err := v.Db.Exec(sql)
	return err
}

func (v *DB) Open(driverName string, dataSourceName string) error {
	v.Db_type = driverName

	var err error

	switch driverName {
	case "sqlite3":
		v.Db, err = sql.Open(driverName, dataSourceName)
	}

	if v.Db == nil {
		err = fmt.Errorf("Unknow driver - %s", driverName)
	}

	return err
}

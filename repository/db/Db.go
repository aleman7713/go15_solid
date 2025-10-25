package db

import (
	"database/sql"
	"fmt"
)

type Idb interface {
	Open(driverName string, dataSourceName string) error
	Close() error
	Exec(sql string) error
}

type DB struct {
	Db_type string
	Db      *sql.DB
}

func (v *DB) Close() error {
	err := v.Db.Close()
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

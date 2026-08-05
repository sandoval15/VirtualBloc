package db

import (
	"database/sql"

	so "github.com/sandoval15/SecureOnce-Go"

	_ "github.com/go-sql-driver/mysql"
)

var (
	instance *sql.DB
	once	so.SecureOnce
)

func GetConnection() (*sql.DB, error) {
	if err := once.Do(func(r *error) {
		instance, *r = sql.Open("mysql", "")
		if *r != nil { return }

		instance.SetMaxOpenConns(2)
		instance.SetMaxIdleConns(2)

		if *r = instance.Ping();  *r != nil {
			instance.Close() 
			return
		}
	}, nil); err != nil {
		return nil, err
	}

	return instance, nil
}
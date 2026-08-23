package db

import (
	"database/sql"
	"os"

	so "github.com/sandoval15/SecureOnce-Go"

	_ "github.com/go-sql-driver/mysql"
)

// dsnLocal se rellena en local y el fichero se marca con skip-worktree para que
// las credenciales no lleguen al repositorio. En produccion la conexion llega
// por la variable de entorno DB_DSN y la constante viaja vacia.
const dsnLocal = ""

func dsn() string {
	if d := os.Getenv("DB_DSN"); d != "" {
		return d
	}

	return dsnLocal
}

var (
	instance *sql.DB
	once     so.SecureOnce
)

func GetConnection() (*sql.DB, error) {
	if err := once.Do(func() error {
		var err error

		if instance, err = sql.Open("mysql", dsn()); err != nil {
			return err
		}

		instance.SetMaxOpenConns(2)
		instance.SetMaxIdleConns(2)

		if err = instance.Ping(); err != nil {
			instance.Close()
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

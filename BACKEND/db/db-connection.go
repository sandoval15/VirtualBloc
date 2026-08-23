package db

import (
	"database/sql"

	so "github.com/sandoval15/SecureOnce-Go"

	_ "github.com/go-sql-driver/mysql"
)

// dsn se rellena en local y el fichero se marca con skip-worktree para que las
// credenciales no lleguen al repositorio.
const dsn = ""

var (
	instance *sql.DB
	once     so.SecureOnce
)

func GetConnection() (*sql.DB, error) {
	if err := once.Do(func() error {
		var err error

		if instance, err = sql.Open("mysql", dsn); err != nil {
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

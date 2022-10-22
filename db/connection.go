package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	connectionString := "abc"
	db, _ := sql.Open("mysql", connectionString)
	return db, nil
}

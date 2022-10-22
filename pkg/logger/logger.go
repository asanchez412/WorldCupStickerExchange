package logger

import "database/sql"

type Logger interface {
}

type logger struct {
	db *sql.DB
}

func NewLogger(db *sql.DB) Logger {
	return &logger{
		db: db,
	}
}

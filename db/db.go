package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

type DB struct {
	db *sql.DB
}

func New() (*DB, error) {
	cfg := mysql.Config{User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() {
	_ = db.db.Close()
}

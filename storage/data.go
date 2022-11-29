package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./storage.db")

	if err != nil {
		log.Fatalf("DB open error: %v", err)
	}

	CreateTable()

	return db.Ping()
}

func CreateTable() {
	CreateProductTable()
	CreateUserTable()
	CreateDailyTable()
}

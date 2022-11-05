package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./storage.db")

	if err != nil {
		log.Fatal(err)
	}

	CreateTable()

	return db.Ping()
}

func CreateTable() {
	CreateWishTable()
	CreateUserTable()
}

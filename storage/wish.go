package storage

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Wish struct {
	Id   int
	Url  string
	Skus []string
}

func CreateWishTable() {
	sqlStmt := `
	create table if not exists wish (
        id integer primary key, 
        url text,
        skus text
    );
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertWish(idString string, url string, skus []string) {
	db, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or replace into wish(id, url, skus) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id, url, strings.Join(skus, ","))
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func LoadWish() []Wish {
	db, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select id, url, skus from wish")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var wishList []Wish

	for rows.Next() {
		var id int
		var url string
		var skus string

		err = rows.Scan(&id, &url, &skus)
		if err != nil {
			log.Fatal(err)
		}

		wishList = append(wishList, Wish{
			Id:   id,
			Url:  url,
			Skus: strings.Split(skus, ","),
		})
	}

	return wishList
}

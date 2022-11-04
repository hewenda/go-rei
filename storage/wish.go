package storage

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Wish struct {
	Url  string
	Skus []string
}

func init() {
	db, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStmt := `
	create table if not exists wish (
        id integer primary key autoincrement, 
        url text,
        skus text
    );
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertWish(url string, skus []string) {
	db, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into wish(url, skus) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(url, strings.Join(skus, ","))
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

	rows, err := db.Query("select url, skus from wish")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var wishList []Wish

	for rows.Next() {
		var url string
		var skus string
		err = rows.Scan(&url, &skus)
		if err != nil {
			log.Fatal(err)
		}

		wishList = append(wishList, Wish{
			Url:  url,
			Skus: strings.Split(skus, ","),
		})
	}

	return wishList
}

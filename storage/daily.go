package storage

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Daily struct {
	Id      int32
	Sku     string
	Compare float64
	Price   float64
}

func CreateDailyTable() {
	sqlStmt := `
	create table if not exists daily (
        id integer primary key autoincrement, 
        sku text,
        compare NUMERIC,
        price NUMERIC
    );
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertDailySku(sku string, compare float64, price float64) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or replace into daily(sku, compare, price) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sku, compare, price)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func ClearDaily() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("delete from daily")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func QueryDailySku() []Daily {
	rows, err := db.Query("select id, sku, compare, price from daily")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var dailyList []Daily

	for rows.Next() {
		var id int32
		var sku string
		var compare float64
		var price float64

		err = rows.Scan(&id, &sku, &compare, &price)
		if err != nil {
			log.Fatal(err)
		}

		dailyList = append(dailyList, Daily{
			Id:      id,
			Sku:     sku,
			Compare: compare,
			Price:   price,
		})
	}

	return dailyList
}

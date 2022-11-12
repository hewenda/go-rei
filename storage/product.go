package storage

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id  int
	Url string
}

func CreateProductTable() {
	sqlStmt := `
	create table if not exists product (
        id integer primary key, 
        url text
    );
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertProduct(idString string, url string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or replace into product(id, url) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id, url)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func QueryProduct() []Product {
	rows, err := db.Query("select id, url from product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var product []Product

	for rows.Next() {
		var id int
		var url string

		err = rows.Scan(&id, &url)
		if err != nil {
			log.Fatal(err)
		}

		product = append(product, Product{
			Id:  id,
			Url: url,
		})
	}

	return product
}

func DeleteProduct(id string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("delete from product where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

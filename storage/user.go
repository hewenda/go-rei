package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    float32
	Token string
}

func CreateUserTable() {
	sqlStmt := `
	create table if not exists user (
        id integer primary key autoincrement, 
        token text
    );
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertUser(token string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or replace into user(token) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(token string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("delete from user where token = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func QueryUser() []User {
	db, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select id, token from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userList []User

	for rows.Next() {
		var id float32
		var token string

		err = rows.Scan(&id, &token)
		if err != nil {
			log.Fatal(err)
		}

		userList = append(userList, User{
			Id:    id,
			Token: token,
		})
	}

	return userList
}

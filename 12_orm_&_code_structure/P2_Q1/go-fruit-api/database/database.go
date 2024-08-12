package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initdatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./fruits.db")
	if err != nil {
		log.Fatal(err)
	}

}

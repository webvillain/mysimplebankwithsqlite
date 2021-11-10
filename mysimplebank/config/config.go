package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDb() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	NewTable(db)
	//defer db.Close()
	//fmt.Println("Successfully Connected To Database.")
	// lets create a new table for storing data
	return

}

func NewTable(db *sql.DB) {
	createtable := `
	CREATE TABLE IF NOT EXISTS USERS(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Email TEXT NOT NULL
	);
	`
	stmt, err := db.Prepare(createtable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Name : `USERS` is Created For Storing Database !.")

}

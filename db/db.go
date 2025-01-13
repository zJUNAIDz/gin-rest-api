package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("failed to Initialize DB")
	}
	DB.SetMaxOpenConns(5)
	DB.SetConnMaxIdleTime(1)
	CreateTable()
}

func CreateTable() {
	createTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userID INTEGER
	)
	`
	_, err := DB.Exec(createTable)
	if err != nil {
		panic("failed to create table" + err.Error())
	}
}

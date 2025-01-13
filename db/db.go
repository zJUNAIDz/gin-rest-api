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
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(userTable)
	if err != nil {
		panic("failed to create users table" + err.Error())
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		panic("failed to create events table" + err.Error())
	}
}

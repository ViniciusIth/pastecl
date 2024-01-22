package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var Access *sql.DB

func ConnectDB() error {
	db, err := sql.Open("sqlite", "file:./db/data.db")
	if err != nil {
		return err
	}

	log.Println("Database started")

	Access = db

	return nil
}

func InitializeDB() error {
	// User table initialization
	// _, err := Access.Exec("CREATE TABLE user (id TEXT NOT NULL PRIMARY KEY, name TEXT NOT NULL, createdAt TEXT NOT NULL)")
	_, err := Access.Exec(`CREATE TABLE IF NOT EXISTS paste (
        id TEXT NOT NULL PRIMARY KEY,
        title TEXT NOT NULL,
        createdat INTEGER NOT NULL,
        expiresat INTEGER NOT NULL,
        visibility INTEGER DEFAULT 1,
        controlkey STRING,
        ownerid STRING
        )`)
	if err != nil {
		return err
	}

	return nil
}

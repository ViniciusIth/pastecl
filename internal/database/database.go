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

	_, err = Access.Exec(`CREATE TABLE IF NOT EXISTS users (
        uuid TEXT NOT NULL PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at INTEGER NOT NULL
        );`)

	if err != nil {
		return err
	}

	return nil
}

package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// this function opens sqlite database
func OpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "database.sqlite3")

	if err != nil {
		return nil
	}

	return db
}

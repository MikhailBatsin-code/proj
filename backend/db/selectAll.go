package db

import (
	"database/sql"
	"fmt"
)

func SelectAll(table string, db *sql.DB) *sql.Rows {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))

	if err != nil {
		return nil
	}

	defer db.Close()

	return rows
}

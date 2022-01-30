package db

import (
	"database/sql"
	"fmt"
)

func SelectById(table string, id string, db *sql.DB) *sql.Row {
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=%s", table, id))

	defer db.Close()

	return row
}

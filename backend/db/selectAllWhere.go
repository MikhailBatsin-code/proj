package db

import (
	"database/sql"
	"fmt"
)

func SelectAllWhere(table string, condition string, db *sql.DB) *sql.Rows {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s;", table, condition))

	if err != nil {
		return nil
	}

	return rows
}

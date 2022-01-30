package db

import "database/sql"

// this function will create tables todo_lists and todos if it has errors it just panics because it only called before runtime
func CreateTables(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS todo_lists(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT UNIQUE);")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT UNIQUE NOT NULL, description TEXT, todo_list INTEGER, FOREIGN KEY(todo_list) REFERENCES todo_lists(id) ON DELETE CASCADE);")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

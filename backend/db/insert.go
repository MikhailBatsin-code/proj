package db

import (
	"database/sql"
	"fmt"

	"github.com/MikhailBatsin-code/proj/backend/models"
)

func InsertIntoTodoLists(tl models.TodoList, db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO todo_lists(title) VALUES(\"%s\");", tl.Title))

	defer db.Close()

	return err
}

func InsertIntoTodos(t models.Todo, db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO todos(title, description, todo_list) VALUES(\"%s\", \"%s\", %s);", t.Title, t.Description, t.TodoListId))

	defer db.Close()

	return err
}

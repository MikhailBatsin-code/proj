package handlers

import (
	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/MikhailBatsin-code/proj/backend/models"
	"github.com/gin-gonic/gin"
)

var GATITLPath string = "/api/v1/todos/:todo_list_id"

func GetAllTodosByTodoList(c *gin.Context) {
	todoLists := db.GetAllTodoLists(c)
	todos := make([]models.Todo, 0)
	todo := models.Todo{}

	if len(todoLists) == 0 {
		c.IndentedJSON(400, gin.H{"status": "there is no todo lists"})
	}

	rows := db.SelectAllWhere("todos", "todo_list="+c.Param("todo_list_id"), db.OpenDb())

	if rows == nil {
		c.IndentedJSON(500, gin.H{"status": "can't query from db"})

		return
	}

	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.TodoListId); err != nil {
			c.IndentedJSON(500, gin.H{"status": "can't scan query value"})

			break
		}

		todos = append(todos, todo)
	}

	c.IndentedJSON(200, todos)
}

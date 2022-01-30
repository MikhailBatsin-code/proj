package handlers

import (
	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/MikhailBatsin-code/proj/backend/models"
	"github.com/gin-gonic/gin"
)

var GTLBIPath string = "/api/v1/todo-lists/:id"

func GetTodoListById(c *gin.Context) {
	var todoList models.TodoList = models.TodoList{}

	row := db.SelectById("todo_lists", c.Param("id"), db.OpenDb())

	if err := row.Scan(&todoList.ID, &todoList.Title); err != nil {
		c.IndentedJSON(500, map[string]string{
			"status": "can't scan query result",
		})

		return
	}

	c.IndentedJSON(200, todoList)
}

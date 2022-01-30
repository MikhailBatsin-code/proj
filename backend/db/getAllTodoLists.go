package db

import (
	"github.com/MikhailBatsin-code/proj/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllTodoLists(c *gin.Context) []models.TodoList {
	var todoLists []models.TodoList = make([]models.TodoList, 0)
	var todoList models.TodoList

	rows := SelectAll("todo_lists", OpenDb())

	if rows == nil {
		c.IndentedJSON(404, map[string]string{
			"status": "can't query todo_lists table",
		})
	}

	for rows.Next() {
		err := rows.Scan(&todoList.ID, &todoList.Title)

		if err != nil {
			c.IndentedJSON(500, map[string]string{
				"status": "can't scan queryed rows",
			})

			break
		}

		todoLists = append(todoLists, todoList)
	}

	return todoLists
}

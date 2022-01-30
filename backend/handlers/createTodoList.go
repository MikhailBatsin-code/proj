package handlers

import (
	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/MikhailBatsin-code/proj/backend/models"
	"github.com/gin-gonic/gin"
)

var CreateTodoListPath string = "/api/v1/todo-lists/create"

func CreateTodoList(c *gin.Context) {
	var todoList models.TodoList

	err := c.BindJSON(&todoList)

	if err != nil {
		c.IndentedJSON(400, gin.H{"status": "can't bind request"})

		return
	}

	err = db.InsertIntoTodoLists(todoList, db.OpenDb())

	if err != nil {
		c.IndentedJSON(500, gin.H{"status": "can't insert into db"})

		return
	}

	c.IndentedJSON(200, gin.H{"status": "created"})
}

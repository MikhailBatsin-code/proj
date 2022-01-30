package handlers

import (
	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/MikhailBatsin-code/proj/backend/models"
	"github.com/gin-gonic/gin"
)

var CreateTodoPath string = "/api/v1/todos/create"

func CreateTodo(c *gin.Context) {
	todo := models.Todo{}

	if err := c.BindJSON(&todo); err != nil {
		c.IndentedJSON(400, gin.H{"status": "invalid json"})
	}

	err := db.InsertIntoTodos(todo, db.OpenDb())

	if err != nil {
		c.IndentedJSON(500, gin.H{"status": "can't insert into db"})
	}

	c.IndentedJSON(200, gin.H{"status": "ok"})
}

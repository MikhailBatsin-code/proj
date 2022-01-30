package handlers

import (
	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/gin-gonic/gin"
)

var GATLPath = "/api/v1/todo-lists"

func GetAllTodoLists(c *gin.Context) {
	todoLists := db.GetAllTodoLists(c)

	c.IndentedJSON(200, todoLists)
}

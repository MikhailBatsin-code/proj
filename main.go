package main

import (
	"log"

	"github.com/MikhailBatsin-code/proj/backend/db"
	"github.com/MikhailBatsin-code/proj/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// create tables
	db.CreateTables(db.OpenDb())

	// initialize gin router
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.1"})

	// routes here
	// GET
	r.GET(handlers.IndexPath, handlers.IndexHandler)
	r.GET(handlers.GATLPath, handlers.GetAllTodoLists)
	r.GET(handlers.GTLBIPath, handlers.GetTodoListById)
	r.GET(handlers.GATITLPath, handlers.GetAllTodosByTodoList)

	// POST
	r.POST(handlers.CreateTodoListPath, handlers.CreateTodoList)
	r.POST(handlers.CreateTodoPath, handlers.CreateTodo)

	// start app and if it has errors print them
	log.Fatal(r.Run(":6969"))
}

package handlers

import "github.com/gin-gonic/gin"

// just index handler needed for testing
var IndexPath string = "/api"

func IndexHandler(c *gin.Context) {
	c.IndentedJSON(200, map[string]string{
		"status": "ok",
	})
}

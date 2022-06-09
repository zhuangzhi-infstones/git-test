package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	server := gin.Default()
	server.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	return server
}

func main() {
	server := Server()
	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

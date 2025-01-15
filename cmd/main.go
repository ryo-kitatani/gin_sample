package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	return r
}

func main() {
	log.Println("Starting server on :8080")
	r := setupRouter()
	r.Run(":8080")
}

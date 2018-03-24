package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//curl "127.0.0.1:8080"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin\n")
	})
	router.Run(":8080")
}

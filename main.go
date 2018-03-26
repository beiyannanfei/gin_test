package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	//curl "127.0.0.1:8080"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin\n")
	})

	//curl 127.0.0.1:8080/user/abcd
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Printf("===== params name: %v\n", name)
		c.String(http.StatusOK, "Hello %s\n", name)
	})

	router.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "world!")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, " + name,
		})
	})

	_ = r.Run(":8082")
}

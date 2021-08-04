package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	cfg := cors.DefaultConfig()
	cfg.AllowWildcard = true
	cfg.AllowOrigins = []string{"http://*.tiket.com", "https://*.tiket.com"}
	cfg.AddAllowHeaders("Authorization", "Content-Type")
	r.Use(cors.New(cfg))

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "world!")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, " + name,
		})
	})

	_ = r.Run(":8082")
}

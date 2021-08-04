package handlers

import "github.com/gin-gonic/gin"

type Context struct {
	Gin *gin.Engine
}

func(c *Context) RegisterRoutes() {
	c.Gin.GET("/people", c.allPeople)
	c.Gin.GET("/people/:peopleId", c.getPerson)
	c.Gin.POST("/people", c.addNewPerson)
	c.Gin.DELETE("/people/:peopleId", c.deletePerson)
}

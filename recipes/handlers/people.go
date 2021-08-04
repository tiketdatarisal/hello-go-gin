package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type People []Person

var (
	data = People{
		{1, "John Doe", 45},
		{2, "Maria Mercedes", 36},
		{3, "Ali Baba", 40},
		{4, "James", 25},
		{5, "Andi Wahyu", 21},
	}
)

func (c *Context) allPeople(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusOK, data)
		return
	}

	result := People{}
	for _, p := range data {
		if strings.HasPrefix(strings.ToLower(p.Name), strings.ToLower(name)) {
			result = append(result, p)
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Context) getPerson(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("peopleId"))

	for _, p := range data {
		if p.ID == id {
			ctx.JSON(http.StatusOK, p)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, nil)
}

func (c *Context) addNewPerson(ctx *gin.Context) {
	p := Person{}
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := data[len(data)-1].ID + 1
	p.ID = id
	data = append(data, p)

	ctx.JSON(http.StatusCreated, p)
}

func (c *Context) deletePerson(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("peopleId"))

	idx := -1
	for i, p := range data {
		if p.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	data = append(data[:idx], data[idx+1:]...)
	ctx.JSON(http.StatusNoContent, nil)
}

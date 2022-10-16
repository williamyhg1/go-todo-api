package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()
	// r.GET("/todos", getTodos)
	// r.Run()

  r.GET("/todos", func(c *gin.Context) {
    c.JSON(http.StatusOK, todos)
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"
}
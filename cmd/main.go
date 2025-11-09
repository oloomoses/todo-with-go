package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/todo/internal/config"
	"github.com/oloomoses/todo/internal/db"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbconn := db.InitDB()
	db.Migrate(dbconn)

	r := gin.Default()

	r.GET("/", home)
	r.GET("/todos", getTodos)
	r.POST("/todo", createTodo)
	r.PATCH("todo/{id}", updateTodo)
	r.DELETE("/todo/{id}", deleteTodo)

	r.Run()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Pong",
	})

}

func getTodos(c *gin.Context) {

}

func createTodo(c *gin.Context) {

}

func updateTodo(c *gin.Context) {

}

func deleteTodo(c *gin.Context) {

}

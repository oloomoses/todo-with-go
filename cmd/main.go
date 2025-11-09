package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/todo/internal/config"
	"github.com/oloomoses/todo/internal/db"
	"github.com/oloomoses/todo/internal/handler"
	"github.com/oloomoses/todo/internal/repository"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbconn := db.InitDB()
	db.Migrate(dbconn)

	r := gin.Default()

	newTodo := repository.NewTodoRepo(dbconn)
	todo := handler.NewTodohHandler(newTodo)

	r.GET("/", home)
	r.GET("/todos", todo.All)
	r.POST("/todo", todo.Create)
	r.GET("/todo/:id", todo.Find)
	r.PATCH("/todo/{id}", todo.Update)
	r.DELETE("/todo/{id}", todo.Delete)

	r.Run()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Pong",
	})

}

func getTodos(c *gin.Context) {

}

func updateTodo(c *gin.Context) {

}

func deleteTodo(c *gin.Context) {

}

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
	todoRepo := repository.NewTodoRepo(dbconn)
	todo := handler.NewTodoHandler(todoRepo)
	// db.Migrate(dbconn)

	r := gin.Default()

	r.GET("/todos", todo.All)
	r.GET("/todo/:id", todo.Find)
	r.POST("/todo", todo.Create)
	r.PUT("/todo/:id", todo.Update)
	r.DELETE("/todo/:id", todo.Delete)

	r.Run()
}

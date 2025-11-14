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

	r := gin.Default()

	// tmpl := template.Must(template.ParseGlob(filepath.Join("templates/**/*.html")))
	// r.SetHTMLTemplate(tmpl)

	r.LoadHTMLGlob("templates/**/*.html")

	r.Static("/static", "./static")

	dbconn := db.InitDB()
	todoRepo := repository.NewTodoRepo(dbconn)
	todo := handler.NewTodoHandler(todoRepo)

	// todo web handler
	todoWeb := handler.NewTodoWebHandler(todoRepo)

	// db.Migrate(dbconn)

	v1 := r.Group("/api/v1")

	{
		v1.GET("/todos", todo.All)
		v1.GET("/todo/:id", todo.Find)
		v1.POST("/todo", todo.Create)
		v1.PUT("/todo/:id", todo.Update)
		v1.DELETE("/todo/:id", todo.Delete)
	}

	r.GET("/", todoWeb.Index)
	r.GET("/todos/:id", todoWeb.Show)

	r.Run()
}

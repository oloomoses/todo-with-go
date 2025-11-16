package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/todo/internal/config"
	"github.com/oloomoses/todo/internal/db"
	"github.com/oloomoses/todo/internal/handler"
	"github.com/oloomoses/todo/internal/repository"
)

func MethodOveride() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			method := c.PostForm("_method")

			if method != "" {
				c.Request.Method = method
			}
		}
		c.Next()
	}
}

func main() {
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	r := gin.Default()
	r.Use(MethodOveride())

	// pattern := filepath.Join("templates", "**", "*.html")
	// tmpl := template.Must(template.ParseGlob(pattern))
	// r.SetHTMLTemplate(tmpl)

	r.LoadHTMLGlob("templates/**/*")

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
	r.GET("/todo/new", todoWeb.NewTodoForm)
	r.POST("/todo", todoWeb.New)
	r.GET("todos/:id/edit", todoWeb.Edit)
	r.POST("todos/:id", todoWeb.Update)

	r.Run()
}

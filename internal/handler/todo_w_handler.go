package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	todo "github.com/oloomoses/todo/internal/model"
	"github.com/oloomoses/todo/internal/repository"
)

type TodoWebHandler struct {
	repo *repository.TodoRepo
}

func NewTodoWebHandler(repo *repository.TodoRepo) *TodoWebHandler {
	return &TodoWebHandler{repo: repo}
}

func (h *TodoWebHandler) Index(c *gin.Context) {
	todos, err := h.repo.FindAll()

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading Todos")
		return
	}

	c.HTML(http.StatusOK, "todos/index.html", gin.H{"Todos": todos})
}

func (h *TodoWebHandler) Show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	todo, err := h.repo.FindByID(uint64(id))

	if err != nil {
		c.String(http.StatusNotFound, "Todo item not found!")
		return
	}

	c.HTML(http.StatusOK, "todos/show.html", gin.H{"Todo": todo})
}

func (h *TodoWebHandler) NewTodoForm(c *gin.Context) {
	c.HTML(http.StatusOK, "todos/new.html", gin.H{
		"Title":     "",
		"Completed": false,
	})
}

func (h *TodoWebHandler) New(c *gin.Context) {
	title := strings.TrimSpace(c.PostForm("title"))
	completedstr := c.PostForm("completed")

	completed := completedstr == "on"

	if title == "" {
		c.HTML(http.StatusBadRequest, "todos/new.html", gin.H{
			"Error":     "Title cannot be blank",
			"Title":     title,
			"Completed": completed,
		})
		return
	}

	input := todo.Todo{
		Title:     title,
		Completed: completed,
	}

	if err := h.repo.Create(&input); err != nil {
		c.HTML(http.StatusInternalServerError, "todos/new.html", gin.H{
			"Error":     "Failed to create todo",
			"Title":     title,
			"Completed": completed,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/todos")

}

func (h *TodoWebHandler) Edit(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	todo, err := h.repo.FindByID(id)

	if err != nil {
		c.HTML(http.StatusBadRequest, "todos/edit.html", gin.H{
			"Error": "Todo not Found",
		})
		return
	}

	c.HTML(http.StatusOK, "todos/edit.html", gin.H{"Todo": todo, "ID": id})
}

func (h *TodoWebHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	title := strings.TrimSpace(c.PostForm("title"))
	completed := c.PostForm("completed") == "on"

	if title == "" {
		c.HTML(http.StatusBadRequest, "todos/edit.html", gin.H{
			"Error": "Title cannot be empty",
		})
		return
	}

	updates := make(map[string]interface{})
	updates["Title"] = title
	updates["Completed"] = completed

	if err := h.repo.Update(id, updates); err != nil {
		c.HTML(http.StatusInternalServerError, "todos/edit.html", gin.H{"Error": "Update Failed"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

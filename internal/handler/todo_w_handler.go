package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	c.HTML(http.StatusOK, "todos/index.html", gin.H{"Todo": todos})
}

func (h *TodoWebHandler) Show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	todo, err := h.repo.FindByID(uint64(id))

	if err != nil {
		c.String(http.StatusNotFound, "Todo item not found!")
		return
	}

	c.HTML(http.StatusOK, "/todos/show.html", gin.H{"Todo": todo})
}

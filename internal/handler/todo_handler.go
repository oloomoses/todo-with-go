package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/oloomoses/todo/internal/model"
	"github.com/oloomoses/todo/internal/repository"
)

type TodoHandler struct {
	repo *repository.TodoRepo
}

func NewTodohHandler(repo *repository.TodoRepo) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) Create(c *gin.Context) {

	var input todo.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.repo.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (h *TodoHandler) All(c *gin.Context) {
	todos, err := h.repo.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)

}

func (h *TodoHandler) Find(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	todo, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch todo with the given id"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (r *TodoHandler) Update(c *gin.Context) {

}

func (r *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := r.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Todo with the given id not Deleted!"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

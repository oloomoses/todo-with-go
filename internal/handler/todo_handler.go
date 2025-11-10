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

func NewTodoHandler(repo *repository.TodoRepo) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var input todo.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := h.repo.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error:": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (h *TodoHandler) All(c *gin.Context) {
	todos, err := h.repo.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)

}

func (h *TodoHandler) Find(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	todo, err := h.repo.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "ID not found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Todo not found!"})
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	var input struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updates := make(map[string]interface{})

	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Completed != nil {
		updates["completed"] = *input.Completed
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No fields to update"})
		return
	}

	if err := h.repo.Update(uint64(id), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Todo not updated"})
		return
	}

	todo, _ := h.repo.FindByID(id)

	c.JSON(http.StatusOK, todo)

}

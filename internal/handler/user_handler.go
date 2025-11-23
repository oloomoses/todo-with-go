package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/todo/internal/model"
	"github.com/oloomoses/todo/internal/repository"
)

type UserHandler struct {
	repo *repository.UserRepo
}

func NewUserHandler(repo *repository.UserRepo) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "users/new", gin.H{
		"username": "",
		"password": "",
	})
}

func (h *UserHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	input := model.User{
		Username: strings.ToLower(username),
		Password: password,
	}

	if err := input.Validate(); err != nil {
		c.HTML(http.StatusBadRequest, "users/new", gin.H{"Error": err})
		return
	}

	if err := h.repo.Create(&input); err != nil {
		c.HTML(http.StatusInternalServerError, "users/new", gin.H{"Error": err})
		return
	}

	c.Redirect(http.StatusSeeOther, "users/index")

}

func (h *UserHandler) AllUsers(c *gin.Context) {
	users, err := h.repo.All()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "users/index", gin.H{"Error": err})
		return
	}

	c.HTML(http.StatusOK, "users/index", gin.H{
		"Users": users,
	})
}

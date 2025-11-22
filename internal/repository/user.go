package repository

import (
	user "github.com/oloomoses/todo/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *user.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepo) All() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error

	return users, err
}

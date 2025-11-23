package repository

import (
	user "github.com/oloomoses/todo/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(newUser *user.User) error {

	hashedPass, err := hashePassword(newUser.Password)

	if err != nil {
		return err
	}
	hashedUser := user.User{
		Username: newUser.Username,
		Password: hashedPass,
	}
	return r.db.Create(&hashedUser).Error
}

func (r *UserRepo) All() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error

	return users, err
}

func hashePassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(byte), err
}

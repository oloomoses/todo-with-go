package model

import (
	"errors"
	"time"
	"unicode"
)

type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Validate() error {
	if u.Username == "" || u.Password == "" {
		return errors.New("username or password cannot be blank")
	}

	if len(u.Username) < 3 || len(u.Username) > 10 {
		return errors.New("username must be between 3 and 10 characters")
	}

	hasUpper := false

	for _, r := range u.Password {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}

	if !hasUpper {
		return errors.New("password mu contain atleast on uppercase letter")
	}

	return nil
}

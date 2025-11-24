package auth

import (
	"errors"

	"github.com/google/uuid"
)

var sessionStore = map[string]string{}

func GenerateSession(username string) string {
	sessionId := uuid.New().String()

	sessionStore[sessionId] = username

	return sessionId
}

func VeriFySession(cookieString string) (string, error) {
	username, ok := sessionStore[cookieString]

	if !ok {
		return "", errors.New("Unauthorized")
	}

	return username, nil
}

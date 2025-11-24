package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/todo/internal/service/auth"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieStr, err := c.Cookie("session_id")

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := auth.VeriFySession(cookieStr)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("username", user)
		c.Next()
	}
}

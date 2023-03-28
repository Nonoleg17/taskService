package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskService/internal/usecase"
)

func AuthUser(u usecase.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		if err := u.Check(cookie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

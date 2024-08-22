package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists || user == nil {
			c.IndentedJSON(http.StatusUnauthorized, "error on login")
			c.Abort()
			return
		}
		c.Next()
	}
}

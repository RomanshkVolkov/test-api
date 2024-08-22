package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/gin-gonic/gin"
)

// middeleware for cors and jwt

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().UTC()
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := repository.ExtractDataByToken(token)

		if err == nil && user.ID != 0 {
			c.Set("user", user)
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		c.Next()

		latency := time.Since(t)
		fmt.Println(latency)

		status := c.Writer.Status()
		fmt.Println(status)

	}
}

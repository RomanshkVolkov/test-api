package middleware

import (
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/gin-gonic/gin"
)

func IPWhiteListSwagger() gin.HandlerFunc {
	return func(c *gin.Context) {
		CurrentIP := c.ClientIP()
		exist, err := repository.SwaggerValidateIPAddress(CurrentIP)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

		if !exist {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

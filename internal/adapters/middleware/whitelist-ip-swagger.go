package middleware

import (
	"fmt"
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/gin-gonic/gin"
)

func IPWhiteListSwagger() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentIP := c.ClientIP()
		origin := c.Request.Header.Get("Origin") // get host
		fmt.Println("host", origin)
		fmt.Println("ip", currentIP)
		repo := repository.GetDBConnection("DB_DSN_DOMAIN_1")
		exist, err := repo.SwaggerValidateIPAddress(currentIP)

		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

		if !exist {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

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
		host := c.Request.Host // get host
		fmt.Println("host", host)
		fmt.Println("ip", currentIP)

		repo := repository.GetDBConnection(host)
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

package http

import (
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/adapters/middleware"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(middleware.Middleware())

	AuthRoutes(r)
	UserRoutes(r)
	MailRoutes(r)

	// path routes
	r.GET("/", func(c *gin.Context) {
		req := c.Request
		var userID uint
		id, exist := c.Get("userID")

		if exist {
			userID = id.(uint)
		}

		c.IndentedJSON(http.StatusOK, domain.APIResponse[any, any]{
			Success: true,
			Message: "Welcome to test-api",
			Data: domain.RequestInfo{
				Host:      req.Host,
				IP:        req.RemoteAddr,
				UserAgent: req.UserAgent(),
				UserID:    userID,
			},
		})
	})
}

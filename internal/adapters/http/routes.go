package http

import (
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/adapters/middleware"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, domain.APIResponse[any, any]{
			Success: false,
			Message: "Page not found",
		})
	})

	r.Use(middleware.Middleware())
	r.Static("/static", "/srv/static")

	AuthRoutes(r)
	UserRoutes(r)
	MailRoutes(r)

	// root route
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

package http

import (
	"github.com/RomanshkVolkov/test-api/internal/adapters/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/profile", handler.GetUserProfile)
	}

}

package http

import (
	"github.com/RomanshkVolkov/test-api/internal/adapters/handler"
	"github.com/gin-gonic/gin"
)

func MailRoutes(r *gin.Engine) {
	mail := r.Group("/mail")
	{
		mail.POST("/test", handler.TestEmail)
	}
}

package http

import (
	"github.com/RomanshkVolkov/test-api/internal/adapters/handler"
	"github.com/RomanshkVolkov/test-api/internal/adapters/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	protect := middleware.Protected
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handler.SignIn)
		auth.POST("/sign-up", handler.SignUp)
		auth.PATCH("/forgot-password", handler.SendEmailWithOTPCode)
		auth.POST("/forgot-password/verify", handler.VerifyForgottenPasswordCode)
		auth.PATCH("/forgot-password/reset", handler.ResetForgottenPassword)
		auth.PUT("/change-password", protect(), handler.ChangePassword)
	}
}

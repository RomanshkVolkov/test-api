package handler

import (
	"fmt"
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/RomanshkVolkov/test-api/internal/core/service"
	"github.com/gin-gonic/gin"
)

// @Summary Just Sign In
// @Description Sign in to the application
// @tags Authentication
// @Produce json
// @Param UserCredentials body domain.SignInRequest true "User credentials"
// @Success 200 {object} domain.APIResponse "Successful sign in"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	request, err := ValidateRequest[domain.SignInRequest](c)
	if err != nil {
		return
	}

	authentication, err := service.SignIn(request.Username, request.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "error on login")
		return
	}

	c.IndentedJSON(http.StatusOK, authentication)
}

// @Summary Just Sign Up
// @tags Authentication
// @Produce json
// @Param UserData body domain.NewUser true "Just the user data"
// @Success 200 {object} domain.APIResponse "Return message"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	request, err := ValidateRequest[domain.NewUser](c)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return
	}

	authentication, err := service.SignUp(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "error on register")
		return
	}

	c.IndentedJSON(http.StatusOK, authentication)
}

// @Summary Send an email with the OTP code
// @Description This endpoint will send an email with the OTP code
// @tags Authentication
// @Produce json
// @Param UserIdentity body domain.PasswordResetRequest true "Requires the username to identify the user"
// @Success 200 {object} domain.APIResponse "Return just a message"
// @Failure 401 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/forgot-password [post]
func SendEmailWithOTPCode(c *gin.Context) {
	request, err := ValidateRequest[domain.PasswordResetRequest](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return
	}

	authentication, err := service.ResetPasswordRequest(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "error on reset password")
		return
	}

	c.IndentedJSON(http.StatusOK, authentication)
}

// @Summary Verify the code is valid
// @Description Returns data about the code
// @tags Authentication
// @Produce json
// @Param UserIdentity body domain.ForgottenPasswordCode true "Require the username and the OTP code"
// @Success 200 {object} domain.APIResponse "Information about the code"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/forgot-password/verify [post]
func VerifyForgottenPasswordCode(c *gin.Context) {
	request, err := ValidateRequest[domain.ForgottenPasswordCode](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return
	}

	authentication, err := service.VerifyForgottenPasswordCode(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "error on reset password")
		return
	}

	c.IndentedJSON(http.StatusOK, authentication)
}

// @Summary Change password with the OTP code
// @Description This endpoint will reset the password of the user with the OTP code
// @tags Authentication
// @Produce json
// @Param NewCredentials body domain.ResetForgottenPassword true "New credentials by OTP"
// @Success 200 {object} domain.APIResponse "Operation information"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/forgot-password/reset [post]
func ResetForgottenPassword(c *gin.Context) {
	request, err := ValidateRequest[domain.ResetForgottenPassword](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return
	}

	res, err := service.ResetForgottenPassword(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "error on reset password")
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// @Summary Change password
// @Description This endpoint will change the password of authenticated the user
// @tags Authentication
// @Produce json
// @Security BearerAuth
// @Param NewPassword body domain.ChangePassword true "New password"
// @Success 200 {object} domain.APIResponse "Operation information"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /auth/change-password [post]
func ChangePassword(c *gin.Context) {
	request, err := ValidateRequest[domain.ChangePassword](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return
	}
	user, _ := c.MustGet("user").(repository.CustomClaims)

	res, err := service.ChangePassword(user, request)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "error on change password")
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

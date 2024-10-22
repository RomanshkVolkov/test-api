package handler

import (
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
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	server := service.GetServer(c)
	authentication, err := server.SignIn(request.Username, request.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on sign in", Es: "error al iniciar sesión"}))
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
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	server := service.GetServer(c)
	authentication, err := server.SignUp(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on sign up", Es: "error al registrarse"}))
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
// @Router /auth/forgot-password [patch]
func SendEmailWithOTPCode(c *gin.Context) {
	request, err := ValidateRequest[domain.PasswordResetRequest](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	server := service.GetServer(c)
	authentication, err := server.ResetPasswordRequest(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on send email", Es: "error al enviar el correo"}))
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
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	server := service.GetServer(c)
	authentication, err := server.VerifyForgottenPasswordCode(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on verify code", Es: "error al verificar el código"}))
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
// @Router /auth/forgot-password/reset [patch]
func ResetForgottenPassword(c *gin.Context) {
	request, err := ValidateRequest[domain.ResetForgottenPassword](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	server := service.GetServer(c)
	res, err := server.ResetForgottenPassword(request)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on reset password", Es: "error al restablecer la contraseña"}))
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
// @Router /auth/change-password [put]
func ChangePassword(c *gin.Context) {
	request, err := ValidateRequest[domain.ChangePassword](c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	user, _ := c.MustGet("user").(repository.CustomClaims)

	server := service.GetServer(c)
	res, err := server.ChangePassword(user, request)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ServerError(err, domain.Message{En: "error on change password", Es: "error al cambiar la contraseña"}))
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

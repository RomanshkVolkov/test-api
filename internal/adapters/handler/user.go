package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Just User Profile by token
// @Description Get user profile by token
// @tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "User profile"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context) {

	user, exists := c.Get("userID")
	fmt.Println("user profile")
	fmt.Printf("%v", user)
	fmt.Println(exists)
	c.IndentedJSON(http.StatusOK, "user profile")
}

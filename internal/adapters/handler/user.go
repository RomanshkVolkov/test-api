package handler

import (
	"fmt"
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/RomanshkVolkov/test-api/internal/core/service"
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

// @Summary Just Users Profiles List
// @Description Get users profiles list
// @tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.UserProfiles "Users profiles"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /users/profiles [get]
func GetUsersProfiles(c *gin.Context) {
	server := service.GetServer(c)

	users, err := server.GetUsersProfiles()
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, ServerError(err, domain.Message{En: "error on get users profiles", Es: "error al obtener perfiles de usuarios"}))
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

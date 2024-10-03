package handler

import (
	"fmt"
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/core/service"
	"github.com/gin-gonic/gin"
)

// @Summary Test email sending
// @Description This endpoint send a test email
// @tags Mail
// @Produce json
// @Security none
// @Success 200 {object} domain.APIResponse "Operation information"
// @Failure 400 {object} string "Unhandled error (report it)"
// @Failure 500 {object} string "Server error (report it)"
// @Router /mail/test [post]
func TestEmail(c *gin.Context) {
	mailOptions := &service.MailOptions{
		To: []string{
			"joseguzmandev@gmail.com",
		},
		Subject: "Test email",
		Body:    "This is a test email",
	}
	done, err := service.SendMail(mailOptions)
	if err != nil || !done {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "error sending email")
		return
	}
	c.IndentedJSON(http.StatusOK, "email sent")
}

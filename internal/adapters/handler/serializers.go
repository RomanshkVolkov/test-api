package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var request T
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "error on request")
		return nil, err
	}
	return &request, nil
}

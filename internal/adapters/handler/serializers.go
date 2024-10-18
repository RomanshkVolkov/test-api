package handler

import (
	"strings"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var RequestError = domain.Message{
	En: "Request error",
	Es: "Error en la solicitud",
}

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var request T
	if err := c.BindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func ServerError(err error, message domain.Message) domain.APIResponse[any, any] {
	response := domain.APIResponse[any, any]{
		Success: false,
		Message: message,
		Data:    nil,
		Error:   err,
	}
	return response
}

func GetSubdomain(c *gin.Context) string {
	host := c.Request.Host
	splitedHost := strings.Split(host, ".")

	if len(splitedHost) > 2 {
		return splitedHost[0]
	}

	return ""
}

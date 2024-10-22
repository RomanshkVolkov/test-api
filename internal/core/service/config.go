package service

import "github.com/gin-gonic/gin"

type Server struct {
	Host string
}

func GetServer(c *gin.Context) *Server {
	return &Server{
		Host: c.Request.Host,
	}
}

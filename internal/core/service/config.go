package service

import "github.com/gin-gonic/gin"

type Server struct {
	Host string
}

func GetServer(c *gin.Context) *Server {
	origin := c.Request.Header.Get("Origin")
	return &Server{
		Host: origin,
	}
}

package main

import (
	"flag"
	"fmt"

	docs "github.com/RomanshkVolkov/test-api/cmd/docs"
	"github.com/RomanshkVolkov/test-api/internal/adapters/http"
	"github.com/RomanshkVolkov/test-api/internal/adapters/middleware"
	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GO API
// @version         3.0
// @description     Created by @RomanshkVolkov.
// @termsOfService  http://swagger.io/terms/
// @contact.name
// @contact.email  jose@guz-studio.dev
// @host 				localhost:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()

	repository.DBConnection()

	fmt.Printf("Starting server...")

	// load services
	app := gin.Default()
	docs.SwaggerInfo.Version = "3.0"

	app.GET("/swagger/*any", middleware.IPWhiteListSwagger(), ginSwagger.WrapHandler(swaggerFiles.Handler))

	http.InitRoutes(app)

	// init routes
	app.Run(":8080")

}

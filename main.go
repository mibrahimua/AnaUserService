package main

import (
	"AnaUserService/config"
	"AnaUserService/controller"
	"AnaUserService/docs"
	"AnaUserService/repository"
	"AnaUserService/service"
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	db := config.GetDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	docs.SwaggerInfo.Title = "Ana Store - Environment: "
	docs.SwaggerInfo.Description = "API for user service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/login", userController.Login)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

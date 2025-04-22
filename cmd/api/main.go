package main

import (
	"github.com/gin-gonic/gin"
	"github.com/julio-o-julio/stacklog-user-service/config"
	"github.com/julio-o-julio/stacklog-user-service/internal/handler"
	"github.com/julio-o-julio/stacklog-user-service/internal/repository"
	"github.com/julio-o-julio/stacklog-user-service/internal/usecase"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	db, err := config.Initialize()
	if err != nil {
		logger.Errorf("| pkg main | Config initialization error: %v", err)
		return
	}

	userRepository := repository.NewUserRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepository)

	userHandler := handler.NewUserHandler(userUseCase)

	// Initialize server
	server := gin.Default()

	// Initialize rotes
	v1 := server.Group("/api/v1")
	{
		v1.GET("/user/:id", userHandler.GetUser)
		v1.POST("/user", userHandler.CreateUser)
		v1.PUT("/user", userHandler.UpdateUser)
		v1.DELETE("/user", userHandler.DeleteUser)
	}

	// Run server
	server.Run(":8080")
}
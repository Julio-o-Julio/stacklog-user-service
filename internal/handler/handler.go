package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/julio-o-julio/stacklog-user-service/internal/usecase"
)

func Initialize() {
	// Initialize handler
	router := gin.Default()

	// Initialize rotes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", usecase.GetUserUseCase)
		v1.POST("/user", usecase.CreateUserUseCase)
		v1.PUT("/user", usecase.UpdateUserUseCase)
		v1.DELETE("/user", usecase.DeleteUserUseCase)
	}

	// Run the server
	router.Run(":8080")
}
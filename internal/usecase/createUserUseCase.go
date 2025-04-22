package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserUseCase(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user",
	})
}
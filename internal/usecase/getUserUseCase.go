package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserUseCase(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user",
	})
}
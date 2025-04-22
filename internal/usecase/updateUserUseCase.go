package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserUseCase(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user",
	})
}
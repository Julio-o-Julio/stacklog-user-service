package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserUseCase(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user",
	})
}
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julio-o-julio/stacklog-user-service/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return UserHandler{
		userUsecase: useCase,
	}
}

func (u *UserHandler) GetUser(ctx *gin.Context) {
	user, err := u.userUsecase.GetUser(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user",
	})
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user",
	})
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user",
	})
}
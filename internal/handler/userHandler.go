package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julio-o-julio/stacklog-user-service/internal/domain"
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

func (u *UserHandler) GetUserById(ctx *gin.Context) {
	user, err := u.userUsecase.GetUserById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := u.userUsecase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
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
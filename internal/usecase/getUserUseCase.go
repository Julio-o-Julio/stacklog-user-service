package usecase

import (
	"github.com/julio-o-julio/stacklog-user-service/internal/domain"
	"github.com/julio-o-julio/stacklog-user-service/internal/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repository,
	}
}

func (uu *UserUseCase) GetUser(id string) (domain.User, error) {
	return uu.repository.GetUser(id)
}
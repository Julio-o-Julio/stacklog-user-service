package usecase

import (
	"time"

	"github.com/julio-o-julio/stacklog-user-service/internal/domain"
	"github.com/julio-o-julio/stacklog-user-service/internal/repository"
	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repository,
	}
}

func (uu *UserUseCase) GetUserById(id string) (domain.User, error) {
	return uu.repository.GetUserById(id)
}

func (uu *UserUseCase) CreateUser(user domain.User) (domain.User, error) {
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
			return domain.User{}, err
	}

	user.ID = uuidv7.New().String()
	user.Password = string(hashedPassword)
	user.UpdatedAt = now
	user.CreatedAt = now

	err = uu.repository.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
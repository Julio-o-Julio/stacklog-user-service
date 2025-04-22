package repository

import (
	"database/sql"

	"github.com/julio-o-julio/stacklog-user-service/config"
	"github.com/julio-o-julio/stacklog-user-service/internal/domain"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return UserRepository{
		database: database,
	}
}

func (ur *UserRepository) GetUser(id string) (domain.User, error) {
	logger := config.GetLogger("getuser_repository")

	query := `SELECT id, username, phone, name, password FROM users WHERE id = ?`

	row := ur.database.QueryRow(query, id)
	
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Phone, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Infof("| pkg repository | User with ID %s not found", id)
			return domain.User{}, err
		}
		logger.Errorf("| pkg repository | Error fetching user: %v", err)
		return domain.User{}, err
	}

	return user, nil
}
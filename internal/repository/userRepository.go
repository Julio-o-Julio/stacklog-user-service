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

func (ur *UserRepository) GetUserById(id string) (domain.User, error) {
	logger := config.GetLogger("getuser_repository")

	query := `SELECT id, username, phone, name, password, updated_at, created_at FROM users WHERE id = $1`

	var user domain.User
	err := ur.database.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Phone,
		&user.Name,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)

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

func (ur *UserRepository) CreateUser(user domain.User) error {
	logger := config.GetLogger("createuser_repository")

	query := `INSERT INTO users (id, username, phone, name, password, updated_at, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := ur.database.Exec(query,
		user.ID,
		user.Username,
		user.Phone,
		user.Name,
		user.Password,
		user.UpdatedAt,
		user.CreatedAt,
	)

	if err != nil {
		logger.Errorf("| pkg repository | Error executing query create user: %v", err)
		return err
	}

	return nil
}

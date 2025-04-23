package domain

import (
	"time"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Phone string `json:"phone"`
	Name *string `json:"name"`
	Password string `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
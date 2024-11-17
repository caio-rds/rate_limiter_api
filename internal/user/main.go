package user

import (
	"gorm.io/gorm"
)

type SQLite struct{ *gorm.DB }

func NewSQLite(db *gorm.DB) *SQLite {
	value := SQLite{
		db,
	}
	return &value
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ResponseUser struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RequestUpdateUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Name     *string `json:"name"`
}

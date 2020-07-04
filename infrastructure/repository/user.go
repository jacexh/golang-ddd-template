package repository

import (
	"context"
	"database/sql"

	"{{.Module}}/domain/user"
)

type (
	UserRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(context.Context, *user.UserEntity) error {
	return nil
}

func (ur *UserRepository) GetUserByID(context.Context, string) (*user.UserEntity, error) {
	return nil, nil
}

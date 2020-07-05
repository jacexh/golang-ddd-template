package persistence

import (
	"context"
	"database/sql"

	"{{.Module}}/domain/user"
	"{{.Module}}/types/do"
)

type (
	UserRepository struct {
		db *sql.DB
	}
)

func convert(entity *user.UserEntity) *do.UserDo {
	return &do.UserDo{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(context.Context, *user.UserEntity) error {
	return nil
}

func (ur *UserRepository) GetUserByID(context.Context, string) (*user.UserEntity, error) {
	return nil, nil
}

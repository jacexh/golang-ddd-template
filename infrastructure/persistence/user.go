package persistence

import (
	"context"
	"database/sql"

	"github.com/jacexh/golang-ddd-template/domain/user"
)

type (
	userRepository struct {
		db *sql.DB
	}
)

func BuildUserRepository(db *sql.DB) user.UserRepository {
	return newUserRepository(db)
}

func newUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) SaveUser(context.Context, *user.UserEntity) error {
	return nil
}

func (ur *userRepository) GetUserByID(context.Context, string) (*user.UserEntity, error) {
	return nil, nil
}

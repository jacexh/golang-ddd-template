package persistence

import (
	"context"

	"github.com/jacexh/golang-ddd-template/domain/user"
	"xorm.io/xorm"
)

type (
	userRepository struct {
		db *xorm.Engine
	}
)

func BuildUserRepository(db *xorm.Engine) user.UserRepository {
	return newUserRepository(db)
}

func newUserRepository(db *xorm.Engine) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) SaveUser(context.Context, *user.UserEntity) error {
	return nil
}

func (ur *userRepository) GetUserByID(context.Context, string) (*user.UserEntity, error) {
	return nil, nil
}

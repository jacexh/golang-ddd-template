package user

import "context"

type UserRepository interface {
	CreateUser(context.Context, *UserEntity) error
	GetUserByID(context.Context, string) (*UserEntity, error)
}

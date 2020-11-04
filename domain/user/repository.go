package user

import "context"

type UserRepository interface {
	SaveUser(context.Context, *UserEntity) error
	GetUserByID(context.Context, string) (*UserEntity, error)
}

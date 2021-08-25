package user

import "context"

type Repository interface {
	SaveUser(context.Context, *UserEntity) error
	GetUserByEmail(context.Context, string) (*UserEntity, error)
}

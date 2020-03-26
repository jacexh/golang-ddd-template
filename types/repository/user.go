package repository

import (
	"context"

	"{{.Module}}/types/entity"
)

// UserRepository user对象仓库接口定义
type UserRepository interface {
	CreateUser(context.Context, *entity.User) error
	GetUserByID(context.Context, string) (*entity.User, error)
}

package persistence

import (
	"{{.Module}}/domain/user"
	"{{.Module}}/types/do"
)

func convertUser(entity *user.UserEntity) *do.UserDo {
	return &do.UserDo{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

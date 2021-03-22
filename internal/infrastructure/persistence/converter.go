package persistence

import (
	"{{.Module}}/internal/domain/user"
	"{{.Module}}/internal/infrastructure/do"
)

func convertUser(entity *user.UserEntity) *do.UserDo {
	return &do.UserDo{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

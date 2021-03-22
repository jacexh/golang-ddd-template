package persistence

import (
	"github.com/jacexh/golang-ddd-template/internal/domain/user"
	"github.com/jacexh/golang-ddd-template/internal/infrastructure/do"
)

func convertUser(entity *user.UserEntity) *do.UserDo {
	return &do.UserDo{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

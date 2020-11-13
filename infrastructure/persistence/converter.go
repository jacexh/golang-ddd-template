package persistence

import (
	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/types/do"
)

func convertUser(entity *user.UserEntity) *do.UserDo {
	return &do.UserDo{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

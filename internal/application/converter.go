package application

import (
	"github.com/jacexh/golang-ddd-template/internal/domain/user"
	"github.com/jacexh/golang-ddd-template/router/dto"
)

func convertUser(user *user.UserEntity) *dto.User {
	return &dto.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

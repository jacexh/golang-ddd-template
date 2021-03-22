package application

import (
	"github.com/jacexh/golang-ddd-template/api/dto"
	"github.com/jacexh/golang-ddd-template/internal/domain/user"
)

func convertUser(user *user.UserEntity) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

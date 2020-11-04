package application

import (
	"{{.Module}}/domain/user"
	"{{.Module}}/types/dto"
)

func convertUser(user *user.UserEntity) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

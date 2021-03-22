package application

import (
	"{{.Module}}/api/dto"
	"{{.Module}}/internal/domain/user"
)

func convertUser(user *user.UserEntity) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

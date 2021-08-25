package application

import (
	"{{.Module}}/internal/domain/user"
	"{{.Module}}/internal/transport/dto"
)

func convertUser(user *user.UserEntity) *dto.User {
	return &dto.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

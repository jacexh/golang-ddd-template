package application

import (
	"context"

	"{{.Module}}/domain/user"
	"{{.Module}}/types/dto"
)

var (
	UserApplication *userApplication
)

type (
	userApplication struct {
		repo user.UserRepository
	}
)

func BuildUserApplication(repo user.UserRepository) {
	UserApplication = &userApplication{
		repo: repo,
	}
}

func convert(user *user.UserEntity) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (ua *userApplication) WithUserRepository(repo user.UserRepository) {
	ua.repo = repo
}

func (ua *userApplication) GetUserByID(ctx context.Context, uid string) (*dto.UserDTO, error) {
	user, err := ua.repo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return convert(user), nil
}

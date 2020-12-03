package application

import (
	"context"

	"{{.Module}}/domain/user"
	"{{.Module}}/logger"
	"{{.Module}}/trace"
	"{{.Module}}/types/dto"
	"go.uber.org/zap"
)

var (
	User *userApplication
)

type (
	userApplication struct {
		repo user.UserRepository
	}

	UserApplication interface {
		GetUserByID(context.Context, string) (*dto.UserDTO, error)
	}
)

// BuildUserApplication create user application instance
func BuildUserApplication(repo user.UserRepository) {
	User = &userApplication{
		repo: repo,
	}
}

// GetUserByID return user data transfer object
func (ua *userApplication) GetUserByID(ctx context.Context, uid string) (*dto.UserDTO, error) {
	u, err := ua.repo.GetUserByID(ctx, uid)
	if err != nil {
		logger.Logger.Error("failed to get user by id", zap.String("user_id", uid), zap.Error(err), trace.MustExtractRequestIndexFromCtxAsField(ctx))
		return nil, err
	}
	return convertUser(u), nil
}

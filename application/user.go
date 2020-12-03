package application

import (
	"context"

	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/trace"
	"github.com/jacexh/golang-ddd-template/types/dto"
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
		logger.Logger.Error("failed to get user by id", zap.String("user_id", uid), zap.Error(err), trace.ExtractRequestIndexFromCtxAsField(ctx))
		return nil, err
	}
	return convertUser(u), nil
}

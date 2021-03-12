package application

import (
	"context"

	"{{.Module}}/application/handler"
	"{{.Module}}/domain/event"
	"{{.Module}}/domain/user"
	"{{.Module}}/logger"
	"{{.Module}}/trace"
	"{{.Module}}/types/dto"
	"go.uber.org/zap"
)

var (
	User UserApplication = (*userApplication)(nil)
)

type (
	userApplication struct {
		repo user.Repository
	}

	UserApplication interface {
		CreateUser(context.Context, *dto.UserDTO) error
	}
)

// BuildUserApplication create user application instance
func BuildUserApplication(repo user.Repository) {
	User = &userApplication{
		repo: repo,
	}

	event.Subscribe(user.EventTypeUserCreated, handler.UserPrinter{})
}

// GetUserByEmail return user data transfer object
func (ua *userApplication) CreateUser(ctx context.Context, dto *dto.UserDTO) error {
	_, err := ua.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		logger.Logger.Error("failed to create user", zap.String("user_id", dto.ID), zap.Error(err), trace.MustExtractRequestIndexFromCtxAsField(ctx))
		return err
	}
	u := user.NewUser(dto.Name, "your_password", dto.Email)
	if err := u.Validate(); err != nil {
		logger.Logger.Error("validation failure", zap.Error(err), trace.MustExtractRequestIndexFromCtxAsField(ctx))
		return err
	}
	err = ua.repo.SaveUser(ctx, u)
	return err
}

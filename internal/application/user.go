package application

import (
	"context"

	"{{.Module}}/internal/application/handler"
	"{{.Module}}/internal/domain/user"
	"{{.Module}}/internal/eventbus"
	"{{.Module}}/internal/logger"
	"{{.Module}}/internal/transport/dto"
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
		CreateUser(context.Context, *dto.User) error
	}
)

// BuildUserApplication create user application instance
func BuildUserApplication(repo user.Repository) {
	User = &userApplication{
		repo: repo,
	}
	eventbus.Subscribe(user.EventTypeUserCreated, handler.UserPrinter{})
}

// CreateUser return user data transfer object
func (ua *userApplication) CreateUser(ctx context.Context, dto *dto.User) error {
	_, err := ua.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		logger.Logger.Error("failed to create user", zap.String("user_id", dto.ID), zap.Error(err), logger.MustExtractTracingIDFromCtx(ctx))
		return err
	}
	u, err := user.NewUser(dto.Name, "your_password", dto.Email)
	if err != nil {
		return err
	}
	if err := u.Validate(); err != nil {
		logger.Logger.Error("validation failure", zap.Error(err), logger.MustExtractTracingIDFromCtx(ctx))
		return err
	}
	err = ua.repo.SaveUser(ctx, u)
	return err
}

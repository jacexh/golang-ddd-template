package application

import (
	"context"

	"github.com/jacexh/golang-ddd-template/application/handler"
	"github.com/jacexh/golang-ddd-template/domain/event"
	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/trace"
	"github.com/jacexh/golang-ddd-template/types/dto"
	"go.uber.org/zap"
)

var (
	User UserApplication = (*userApplication)(nil)
)

type (
	userApplication struct {
		repo user.UserRepository
	}

	UserApplication interface {
		CreateUser(context.Context, *dto.UserDTO) error
	}
)

// BuildUserApplication create user application instance
func BuildUserApplication(repo user.UserRepository) {
	User = &userApplication{
		repo: repo,
	}

	event.Subscribe(user.EventTypeUserCreated, handler.UserPrinter{})
}

// GetUserByID return user data transfer object
func (ua *userApplication) CreateUser(ctx context.Context, dto *dto.UserDTO) error {
	_, err := ua.repo.GetUserByID(ctx, dto.ID)
	if err != nil {
		logger.Logger.Error("failed to create user", zap.String("user_id", dto.ID), zap.Error(err), trace.MustExtractRequestIndexFromCtxAsField(ctx))
		return err
	}
	_ = ua.repo.SaveUser(ctx, user.NewUser(dto.ID, dto.Name, "", dto.Email))
	return nil
}

package handler

import (
	"context"

	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/trace"
	"go.uber.org/zap"
)

type (
	PrintEvent struct{}
)

func (pe PrintEvent) Handle(ctx context.Context, event user.Event) error {
	if event.Type() == "user.created" {
		logger.Logger.Info("created a new user", zap.String("user_name", event.(user.UserCreated).Name), trace.MustExtractRequestIndexFromCtxAsField(ctx))
	}
	return nil
}

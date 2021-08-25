package handler

import (
	"context"

	"{{.Module}}/internal/domain/user"
	"{{.Module}}/internal/eventbus"
	"{{.Module}}/internal/logger"
	"go.uber.org/zap"
)

type (
	UserPrinter struct{}
)

func (up UserPrinter) Handle(ctx context.Context, ev eventbus.DomainEvent) {
	logger.Logger.Info("created a new user", zap.String("user_name", ev.(user.EventUserCreated).Name),
		logger.MustExtractTracingIDFromCtx(ctx))
}

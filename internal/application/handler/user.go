package handler

import (
	"context"

	"github.com/jacexh/golang-ddd-template/internal/domain/user"
	"github.com/jacexh/golang-ddd-template/internal/eventbus"
	"github.com/jacexh/golang-ddd-template/internal/logger"
	"go.uber.org/zap"
)

type (
	UserPrinter struct{}
)

func (up UserPrinter) Handle(ctx context.Context, ev eventbus.DomainEvent) {
	logger.Logger.Info("created a new user", zap.String("user_name", ev.(user.EventUserCreated).Name),
		logger.MustExtractTracingIDFromCtx(ctx))
}

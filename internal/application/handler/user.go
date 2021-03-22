package handler

import (
	"context"

	"github.com/jacexh/golang-ddd-template/internal/domain/event"
	"github.com/jacexh/golang-ddd-template/internal/domain/user"
	"github.com/jacexh/golang-ddd-template/internal/logger"
	"github.com/jacexh/golang-ddd-template/internal/trace"
	"go.uber.org/zap"
)

type (
	UserPrinter struct{}
)

func (up UserPrinter) Handle(ctx context.Context, ev event.DomainEvent) {
	logger.Logger.Info("created a new user", zap.String("user_name", ev.(user.EventUserCreated).Name),
		trace.MustExtractRequestIndexFromCtxAsField(ctx))
}

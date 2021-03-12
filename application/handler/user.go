package handler

import (
	"context"

	"{{.Module}}/domain/event"
	"{{.Module}}/domain/user"
	"{{.Module}}/logger"
	"{{.Module}}/trace"
	"go.uber.org/zap"
)

type (
	UserPrinter struct{}
)

func (up UserPrinter) Handle(ctx context.Context, ev event.DomainEvent) {
	logger.Logger.Info("created a new user", zap.String("user_name", ev.(user.EventUserCreated).Name),
		trace.MustExtractRequestIndexFromCtxAsField(ctx))
}

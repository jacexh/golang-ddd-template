package trace

import (
	"context"

	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"go.uber.org/zap"
)

const (
	ZapKeyRequestIndex = "request-index"
)

func MustExtractRequestIndexFromCtxAsField(ctx context.Context) zap.Field {
	return zap.String(ZapKeyRequestIndex, infection.MustExtract(ctx, CtxKeyRequestIndex).(string))
}

func ExtractRequestIndexFromCtxAsField(ctx context.Context) (zap.Field, error) {
	v, err := infection.Extract(ctx, CtxKeyRequestIndex)
	if err != nil {
		return zap.Field{}, err
	}
	return zap.String(ZapKeyRequestIndex, v.(string)), nil
}

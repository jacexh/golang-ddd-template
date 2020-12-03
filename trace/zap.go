package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"go.uber.org/zap"
)

const (
	ZapKeyRequestIndex = "request-index"
)

func ExtractRequestIndexFromCtxAsField(ctx context.Context) zap.Field {
	return zap.String(ZapKeyRequestIndex, infection.MustExtract(ctx, CtxKeyRequestIndex).(string))
}

func ExtractRequestIndexAsField(c *gin.Context) zap.Field {
	return zap.String(ZapKeyRequestIndex, MustExtractRequestIndex(c))
}

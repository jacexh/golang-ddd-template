package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"{{.Module}}/pkg/infection"
)

const (
	CtxKeyRequestIndex = "request-index"
)

func GenContextWithRequestIndex(c *gin.Context) context.Context {
	index := MustExtractRequestIndex(c)
	ctx := infection.GenContextWithDefaultTimeout()
	return infection.Store(ctx, CtxKeyRequestIndex, index)
}

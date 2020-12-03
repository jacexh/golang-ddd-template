package helper

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"github.com/jacexh/golang-ddd-template/router/middleware"
)

const (
	CtxKeyRequestIndex = "request-index"
)

func GenContextWithRequestIndex(c *gin.Context) context.Context {
	index := middleware.MustExtractRequestIndex(c)
	ctx := infection.GenContextWithDefaultTimeout()
	return infection.Store(ctx, CtxKeyRequestIndex, index)
}

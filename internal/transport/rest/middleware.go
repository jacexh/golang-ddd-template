package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/jacexh/golang-ddd-template/internal/logger"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"go.uber.org/zap"
)

type (
	ChiRequestIDTracer struct{}
)

var (
	_ logger.Tracer = (*ChiRequestIDTracer)(nil)
)

const (
	requestIDKey = "request_id"
)

func InfectContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(infection.GenContextWithDefaultTimeout())

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (t *ChiRequestIDTracer) ExtractTracingIDFromCtx(ctx context.Context) (zap.Field, error) {
	return zap.String(requestIDKey, middleware.GetReqID(ctx)), nil
}

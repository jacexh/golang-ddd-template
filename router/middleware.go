package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"go.uber.org/zap"
)

type (
	ZapLogger struct {
		logger    *zap.Logger
		requestID string
	}
)

var (
	_ middleware.LogEntry = (*ZapLogger)(nil)
)

func GlobalTimeout(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(infection.GenContextWithDefaultTimeout())

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func NewZapLogEntry(logger *zap.Logger, r *http.Request) *ZapLogger {
	entry := &ZapLogger{
		logger: logger,
	}
	entry.requestID = middleware.GetReqID(r.Context())

	logger.Info("request started",
		zap.String("method", r.Method),
		zap.String("uri", r.RequestURI),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("user_agent", r.UserAgent()),
		zap.String("request_id", entry.requestID),
	)
	return entry
}

func (log *ZapLogger) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	log.logger.Info("request complete",
		zap.Int("response_status_code", status),
		zap.Int("response_bytes_length", bytes),
		zap.String("elapsed", elapsed.String()),
		zap.String("request_id", log.requestID),
	)
}

func (log *ZapLogger) Panic(v interface{}, stack []byte) {
	log.logger.Error("broken request",
		zap.Any("panic", v),
		zap.ByteString("stack", stack),
		zap.String("request_id", log.requestID),
	)
}

func ZapLog(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := NewZapLogEntry(logger, r)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			next.ServeHTTP(ww, middleware.WithLogEntry(r, entry))

			entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
		}
		return http.HandlerFunc(fn)
	}
}

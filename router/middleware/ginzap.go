package middleware

import (
	"bytes"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	bpool = &bytesPool{
		pool: sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(nil)
		}},
	}
)

type (
	// hijackWriter 通过代理模式拦截
	hijackWriter struct {
		gin.ResponseWriter
		cache *bytes.Buffer
	}

	bytesPool struct {
		pool sync.Pool
	}
)

func (bp *bytesPool) get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

func (bp *bytesPool) put(cache *bytes.Buffer) {
	cache.Reset()
	bp.pool.Put(cache)
}

func (hw *hijackWriter) Write(b []byte) (int, error) {
	hw.cache.Write(b)
	return hw.ResponseWriter.Write(b)
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginzap(logger *zap.Logger, mergeLog bool, dumpResponse bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 避免被其他中间件修改
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		if !mergeLog {
			logger.Info("received http request",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				ExtractRequestIndexAsZapField(c),
			)
		}

		if dumpResponse {
			cache := bpool.get()
			defer bpool.put(cache)
			c.Writer = &hijackWriter{cache: cache, ResponseWriter: c.Writer}
		}

		c.Next()

		latency := time.Since(start).Milliseconds()
		var respBody []byte

		if dumpResponse {
			if c.Writer.Size() > 0 {
				respBody = c.Writer.(*hijackWriter).cache.Bytes()
			}
		}

		switch {
		case len(c.Errors) > 0:
			logger.Error("got some errors from gin.Context", ExtractRequestIndexAsZapField(c), zap.Strings("errors", c.Errors.Errors()))

		case !mergeLog && !dumpResponse:
			logger.Info("send http response",
				zap.Int("status-code", c.Writer.Status()),
				zap.Int64("latency", latency),
				ExtractRequestIndexAsZapField(c),
			)

		case !mergeLog && dumpResponse:
			logger.Info("send http response",
				zap.Int("status", c.Writer.Status()),
				zap.ByteString("response-body", respBody),
				zap.Int64("latency", latency),
				ExtractRequestIndexAsZapField(c),
			)

		case mergeLog && !dumpResponse:
			logger.Info("handled http request",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Int("status", c.Writer.Status()),
				zap.Int64("latency", latency),
				ExtractRequestIndexAsZapField(c),
			)

		case mergeLog && dumpResponse:
			logger.Info("handled http request",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Int("status", c.Writer.Status()),
				zap.ByteString("response-body", respBody),
				zap.Int64("latency", latency),
				ExtractRequestIndexAsZapField(c),
			)
		}
	}
}

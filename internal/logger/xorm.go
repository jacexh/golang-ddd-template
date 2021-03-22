package logger

import (
	"fmt"

	"github.com/jacexh/golang-ddd-template/internal/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xorm.io/xorm/log"
)

type (
	XormZapLogger struct {
		logger *zap.Logger
		off    bool
		show   bool
	}
)

func NewXormZapLogger(zl *zap.Logger) *XormZapLogger {
	return &XormZapLogger{logger: zl.Named("xorm"), off: false, show: true}
}

func (zl *XormZapLogger) BeforeSQL(_ log.LogContext) {
	return
}

func (zl *XormZapLogger) AfterSQL(ctx log.LogContext) {
	field, err := trace.ExtractRequestIndexFromCtxAsField(ctx.Ctx)

	switch {
	case ctx.Err == nil && err == nil:
		zl.logger.Info(ctx.SQL, zap.Any("args", ctx.Args), zap.Duration("execution_time", ctx.ExecuteTime), field)

	case ctx.Err == nil && err != nil:
		zl.logger.Info(ctx.SQL, zap.Any("args", ctx.Args), zap.Duration("execution_time", ctx.ExecuteTime))

	case ctx.Err != nil && err == nil:
		zl.logger.Info(ctx.SQL, zap.Any("args", ctx.Args), zap.Duration("execution_time", ctx.ExecuteTime), zap.Error(ctx.Err), field)

	case ctx.Err != nil && err != nil:
		zl.logger.Info(ctx.SQL, zap.Any("args", ctx.Args), zap.Duration("execution_time", ctx.ExecuteTime), zap.Error(ctx.Err))
	}
}

func (zl *XormZapLogger) Debugf(format string, v ...interface{}) {
	zl.logger.Debug(fmt.Sprintf(format, v...))
}

func (zl *XormZapLogger) Infof(format string, v ...interface{}) {
	zl.logger.Info(fmt.Sprintf(format, v...))
}

func (zl *XormZapLogger) Warnf(format string, v ...interface{}) {
	zl.logger.Warn(fmt.Sprintf(format, v...))
}

func (zl *XormZapLogger) Errorf(format string, v ...interface{}) {
	zl.logger.Error(fmt.Sprintf(format, v...))
}

func (zl *XormZapLogger) Level() log.LogLevel {
	if zl.off {
		return log.LOG_OFF
	}

	for _, l := range []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel} {
		if zl.logger.Core().Enabled(l) {
			switch l {
			case zapcore.DebugLevel:
				return log.LOG_DEBUG

			case zapcore.InfoLevel:
				return log.LOG_INFO

			case zapcore.WarnLevel:
				return log.LOG_WARNING

			case zapcore.ErrorLevel:
				return log.LOG_ERR
			}
		}
	}
	return log.LOG_UNKNOWN
}

func (zl *XormZapLogger) SetLevel(l log.LogLevel) {
	zl.logger.Warn("cannot change zap logger level after created", zap.String("more_details", "https://github.com/uber-go/zap/issues/591"))
}

func (zl *XormZapLogger) ShowSQL(b ...bool) {
	zl.show = b[0]
}
func (zl *XormZapLogger) IsShowSQL() bool {
	return zl.show
}

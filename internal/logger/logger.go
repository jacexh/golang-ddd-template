package logger

import (
	"{{.Module}}/internal/option"
	"github.com/jacexh/gopkg/zaprotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger 全局日志对象
	Logger *zap.Logger

	levelMapper = map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
)

// BuildLogger 构建全局日志
func BuildLogger(opt option.LoggerOption) *zap.Logger {
	conf := zap.NewProductionConfig()
	conf.Sampling = nil
	conf.EncoderConfig.TimeKey = "@timestamp"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.Level = zap.NewAtomicLevelAt(levelMapper[opt.Level])
	Logger = zaprotate.BuildRotateLogger(conf, zaprotate.RotatingFileConfig{
		LoggerName: opt.Name,
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
		LocalTime:  opt.LocalTime,
		Compress:   opt.Compress,
	})
	return Logger
}

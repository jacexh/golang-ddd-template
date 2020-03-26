package logger

import (
	"sync"

	"{{.Module}}/types"
	"github.com/jacexh/goutil/zaphelper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger 全局日志对象
	Logger *zap.Logger

	once        sync.Once
	levelMapper = map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
)

// BuildLogger 构建全局日志
func BuildLogger(opt types.LoggerOption) *zap.Logger {
	once.Do(func() {
		conf := zap.NewProductionConfig()
		conf.Sampling = nil
		conf.EncoderConfig.TimeKey = "@timestamp"
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.Level = zap.NewAtomicLevelAt(levelMapper[opt.Level])
		Logger = zaphelper.BuildRotateLogger(conf, zaphelper.RotatingFileConfig{
			LoggerName: opt.Name,
			Filename:   opt.Filename,
			MaxSize:    opt.MaxSize,
			MaxAge:     opt.MaxAge,
			MaxBackups: opt.MaxBackups,
			LocalTime:  opt.LocalTime,
			Compress:   opt.Compress,
		})
	})
	return Logger
}

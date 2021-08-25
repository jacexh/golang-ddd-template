package logger

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type Tracer interface {
	ExtractTracingIDFromCtx(context.Context) (zap.Field, error)
}

var tracer Tracer

func ExtractTracingIDFromCtx(ctx context.Context) (zap.Field, error) {
	if tracer != nil {
		return tracer.ExtractTracingIDFromCtx(ctx)
	}
	return zap.Field{}, errors.New("no tracer provided")
}

func MustExtractTracingIDFromCtx(ctx context.Context) zap.Field {
	field, err := ExtractTracingIDFromCtx(ctx)
	if err != nil {
		panic(err)
	}
	return field
}

func SetTracer(t Tracer) {
	tracer = t
}

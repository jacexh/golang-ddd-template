package infection

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	KeyValues = "infection-values"
)

var (
	DefaultTimeout = 30 * time.Second
)

var (
	parent context.Context
	cancel context.CancelFunc
)

type (
	Any map[string]interface{}
)

func GenContextWithDefaultTimeout() context.Context {
	c, _ := context.WithTimeout(parent, DefaultTimeout)
	return c
}

func GenContextWithTimeout(d time.Duration) context.Context {
	c, _ := context.WithTimeout(parent, d)
	return c
}

func GenContextWithValues(v Any) context.Context {
	return context.WithValue(GenContextWithDefaultTimeout(), KeyValues, v)
}

func Store(ctx context.Context, key string, value interface{}) context.Context {
	v, exists := ctx.Value(KeyValues).(Any)
	if !exists {
		return context.WithValue(ctx, KeyValues, Any{key: value})
	}
	v[key] = value
	return ctx
}

func StoreAny(ctx context.Context, any Any) context.Context {
	v, exists := ctx.Value(KeyValues).(Any)
	if !exists {
		return context.WithValue(ctx, KeyValues, any)
	}
	for key, value := range any {
		v[key] = value
	}
	return ctx
}

func Extract(ctx context.Context, key string) (interface{}, error) {
	any, ok := ctx.Value(KeyValues).(Any)
	if !ok {
		return nil, errors.New("no value in context")
	}
	v, ok := any[key]
	if !ok {
		return nil, fmt.Errorf("cannot found anything by key: %s", key)
	}
	return v, nil
}

func MustExtract(ctx context.Context, key string) interface{} {
	v, err := Extract(ctx, key)
	if err != nil {
		panic(err)
	}
	return v
}

func KillContextsAfter(d time.Duration) {
	time.Sleep(d)
	cancel()
}

func KillContextsImmediately() {
	cancel()
}

func init() {
	parent, cancel = context.WithCancel(context.Background())
}

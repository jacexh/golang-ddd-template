package middleware

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	RequestIndexer struct {
		name    string
		counter uint64
	}
)

const (
	GinKeyRequestIndex = "request-index"
)

func NewRequestIndexer(name string) *RequestIndexer {
	ri := &RequestIndexer{name: name}
	if ri.name == "" {
		ri.name = "request"
	}
	return ri
}

func (ri *RequestIndexer) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		index := atomic.AddUint64(&ri.counter, 1)
		c.Set(GinKeyRequestIndex, ri.name+"-"+strconv.FormatUint(index, 10))

		c.Next()
	}
}

func ExtractRequestIndex(c *gin.Context) (string, error) {
	i, ok := c.Get(GinKeyRequestIndex)
	if !ok {
		return "", fmt.Errorf("cannot found request index from gin.Context by key: %s", GinKeyRequestIndex)
	}
	index, ok := i.(string)
	if !ok {
		return "", fmt.Errorf("cannot convert %v as string, %T", i, i)
	}
	return index, nil
}

func MustExtractRequestIndex(c *gin.Context) string {
	index, err := ExtractRequestIndex(c)
	if err != nil {
		panic(err)
	}
	return index
}

func ExtractRequestIndexAsZapField(c *gin.Context) zap.Field {
	return zap.String(GinKeyRequestIndex, MustExtractRequestIndex(c))
}

package middleware

import (
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/trace"
)

type (
	RequestIndexer struct {
		name    string
		counter uint64
	}
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
		c.Set(trace.GinKeyRequestIndex, ri.name+"-"+strconv.FormatUint(index, 10))

		c.Next()
	}
}

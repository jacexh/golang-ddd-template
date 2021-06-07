package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/internal/trace"
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
		ri.genName()
	}
	return ri
}

func (ri *RequestIndexer) genName() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	ri.name = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

func (ri *RequestIndexer) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		index := atomic.AddUint64(&ri.counter, 1)
		c.Set(trace.GinKeyRequestIndex, ri.name+"-"+strconv.FormatUint(index, 10))

		c.Next()
	}
}

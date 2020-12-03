package bytespool

import (
	"bytes"
	"sync"
)

type (
	BytesPool struct {
		pool sync.Pool
	}
)

func NewBytesPool() *BytesPool {
	return &BytesPool{pool: sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		}}}
}

func (bp *BytesPool) Get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

func (bp *BytesPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}

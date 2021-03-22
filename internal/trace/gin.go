package trace

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	GinKeyRequestIndex = "request-index"
)

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

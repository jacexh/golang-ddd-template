package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/types"
	"github.com/jacexh/goutil/gin-middleware/ginzap"
)

// BuildRouter 构造Router
func BuildRouter(option types.RouterOption) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		ginzap.RecoveryWithZap(logger.Logger, option.LogStackIfPanic),
		ginzap.Ginzap(logger.Logger, option.MergeLog, option.DumpResponse),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEPlain, []byte("pong!"))
	})

	return router
}

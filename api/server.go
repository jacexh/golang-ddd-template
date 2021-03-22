package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/api/middleware"
	v1 "{{.Module}}/api/v1"
	"{{.Module}}/internal/logger"
	"{{.Module}}/internal/option"
)

// BuildRouter 构造Router
func BuildRouter(option option.RouterOption) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middleware.NewRequestIndexer("").Handle(),
		middleware.RecoveryWithZap(logger.Logger, option.LogStackIfPanic),
		middleware.Ginzap(logger.Logger, option.MergeLog, option.DumpResponse),
	)

	group := router.Group("/v1")
	{
		group.POST("/user", v1.CreateUser)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEPlain, []byte("ping"))
	})
	return router
}

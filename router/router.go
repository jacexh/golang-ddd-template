package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/logger"
	"{{.Module}}/option"
	"{{.Module}}/router/api"
	"{{.Module}}/router/middleware"
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

	group := router.Group("/api")
	{
		group.POST("/user", api.CreateUser)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEPlain, []byte("ping"))
	})
	return router
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/logger"
	"{{.Module}}/option"
	"{{.Module}}/router/api"
	"github.com/jacexh/goutil/gin-middleware/ginzap"
)

// BuildRouter 构造Router
func BuildRouter(option option.RouterOption) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		ginzap.RecoveryWithZap(logger.Logger, option.LogStackIfPanic),
		ginzap.Ginzap(logger.Logger, option.MergeLog, option.DumpResponse),
	)

	group := router.Group("/api")
	{
		group.GET("/users/:user", api.GetUser)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEPlain, []byte("ping"))
	})
	return router
}

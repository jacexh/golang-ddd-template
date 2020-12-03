package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/option"
	"github.com/jacexh/golang-ddd-template/router/api"
	"github.com/jacexh/golang-ddd-template/router/middleware"
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
		group.GET("/users/:user", api.GetUser)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEPlain, []byte("ping"))
	})
	return router
}

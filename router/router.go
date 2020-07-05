package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/application"
	"{{.Module}}/logger"
	"{{.Module}}/option"
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
	return router
}

func Ping(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEPlain, []byte("ping"))
}

func GetUser(c *gin.Context) {
	uid, ok := c.Params.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	dto, err := application.UserApplication.GetUserByID(context.Background(), uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dto)
}

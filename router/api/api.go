package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/application"
)

func GetUser(c *gin.Context) {
	uid := c.Param("user")
	if uid == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	dto, err := application.User.GetUserByID(context.Background(), uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dto)
}

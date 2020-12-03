package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/application"
	"github.com/jacexh/golang-ddd-template/router/helper"
)

func GetUser(c *gin.Context) {
	uid := c.Param("user")
	if uid == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	dto, err := application.User.GetUserByID(helper.GenContextWithRequestIndex(c), uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, dto)
}

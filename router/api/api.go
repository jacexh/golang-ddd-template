package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-ddd-template/application"
	"github.com/jacexh/golang-ddd-template/trace"
	"github.com/jacexh/golang-ddd-template/types/dto"
)

func SaveUser(c *gin.Context) {
	user := new(dto.UserDTO)
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_ = application.User.CreateUser(trace.GenContextWithRequestIndex(c), user)
	c.JSON(http.StatusOK, nil)
}

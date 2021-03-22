package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/api/dto"
	"{{.Module}}/internal/application"
	"{{.Module}}/internal/trace"
)

func CreateUser(c *gin.Context) {
	user := new(dto.UserDTO)
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_ = application.User.CreateUser(trace.GenContextWithRequestIndex(c), user)
	c.JSON(http.StatusOK, nil)
}
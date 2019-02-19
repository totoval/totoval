package groups

import (
	"github.com/gin-gonic/gin"
	"totoval/app/http/controllers"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Register(group *gin.RouterGroup) {
	newGroup := group.Group("/user")
	{
		newGroup.GET("/info", ug.UserController.Info)
	}
}

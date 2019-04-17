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

		newGroup.GET("/update", ug.UserController.Update)
		newGroup.GET("/delete", ug.UserController.Delete)
		newGroup.GET("/delete-transaction", ug.UserController.DeleteTransaction)
		newGroup.GET("/logout", ug.UserController.LogOut)
		newGroup.GET("/restore", ug.UserController.Restore)
	}
}

package groups

import (
	"github.com/totoval/framework/route"
	"totoval/app/http/controllers"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Group(group route.Grouper) {
	group.GET("/info", ug.UserController.Info)

	group.GET("/update", ug.UserController.Update)
	group.GET("/delete", ug.UserController.Delete)
	group.GET("/delete-transaction", ug.UserController.DeleteTransaction)
	group.GET("/logout", ug.UserController.LogOut)
	group.GET("/restore", ug.UserController.Restore)
}

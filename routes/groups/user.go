package groups

import (
	"github.com/totoval/framework/policy"
	"github.com/totoval/framework/route"
	"totoval/app/http/controllers"
	"totoval/app/policies"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Group(group route.Grouper) {
	group.GET("/info/:userId", ug.UserController.Info).Can(policies.NewUserPolicy(), policy.ActionView)

	group.GET("/update", ug.UserController.Update)
	group.GET("/delete", ug.UserController.Delete)
	group.GET("/delete-transaction", ug.UserController.DeleteTransaction)
	group.GET("/logout", ug.UserController.LogOut)
	group.GET("/restore", ug.UserController.Restore)
}

package groups

import (
	"github.com/totoval/framework/route"
	"totoval/app/http/controllers"
)

type AuthGroup struct {
	LoginController    controllers.Login
	RegisterController controllers.Register
}

func (ag *AuthGroup) Group(group route.Grouper) {
	group.POST("/login", ag.LoginController.Login)
	group.POST("/register", ag.RegisterController.Register)
}

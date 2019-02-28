package groups

import (
	"github.com/gin-gonic/gin"
	"totoval/app/http/controllers"
)

type AuthGroup struct {
	LoginController    controllers.Login
	RegisterController controllers.Register
}

func (ag *AuthGroup) Register(group *gin.RouterGroup) {
	newGroup := group.Group("")
	{
		newGroup.POST("/login", ag.LoginController.Login)
		newGroup.POST("/register", ag.RegisterController.Register)

	}
}

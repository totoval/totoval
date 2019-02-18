package groups

import (
	"github.com/gin-gonic/gin"
	"totoval/app/http/controllers"
)

type AuthGroup struct {
	LoginController    *controllers.Login
	RegisterController *controllers.Register
}

func (g *AuthGroup) Register(group *gin.RouterGroup) {
	newGroup := group.Group("")
	{
		newGroup.POST("/login", g.LoginController.Login)
		newGroup.POST("/register", g.RegisterController.Register)
	}
}

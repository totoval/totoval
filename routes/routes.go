package routes

import (
	"github.com/gin-gonic/gin"
	"totoval/app/http/controllers"
)

type RouteGrouper interface {
	Register(group *gin.RouterGroup)
}

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

type UserGroup struct {
	UserController *controllers.User
}

func (g *UserGroup) Register(group *gin.RouterGroup) {
	newGroup := group.Group("/user")
	{
		newGroup.GET("/info", g.UserController.Info)
	}
}

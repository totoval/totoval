package routes

import (
	"Wallet/app/http/controllers"
	"github.com/gin-gonic/gin"
)

type Auth struct {}
func (Auth) Register (group *gin.RouterGroup){
	newGroup := group.Group("")
	{
		newGroup.POST("/login", controllers.Login{}.Login)
		newGroup.POST("/register", controllers.Register{}.Register)
	}
}


type User struct {}
func (User) Register (group *gin.RouterGroup){
	newGroup := group.Group("/user")
	{
		newGroup.GET("/info", controllers.User{}.Info)
	}
}
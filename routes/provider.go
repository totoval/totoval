package routes

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Router *gin.Engine
}

func (routes Routes) Register(){
	routes.v1()
}

func (routes Routes) v1(){
	v1 := routes.Router.Group("/v1")
	{
		noAuth(v1)
		auth(v1)
	}
}

func noAuth(group *gin.RouterGroup){
	Auth{}.Register(group)
}

func auth(group *gin.RouterGroup){
	User{}.Register(group)
}


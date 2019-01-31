package routes

import (
	"github.com/gin-gonic/gin"
	"totoval/app/http/middleware"
)

type Routes struct {
	Router *gin.Engine
}

func (routes *Routes) Register() {
	routes.v1()
}

func (routes *Routes) v1() {
	v1 := routes.Router.Group("/v1")
	{
		noAuth(v1)
		auth(v1)
	}
}

func registerRouteGroup(g RouteGrouper, group *gin.RouterGroup) {
	g.Register(group)
}

func noAuth(group *gin.RouterGroup) {
	registerRouteGroup(&AuthGroup{}, group)
}

func auth(group *gin.RouterGroup) {
	authGroup := group.Group("", middleware.AuthRequired())

	{
		registerRouteGroup(&UserGroup{}, authGroup)
	}
}

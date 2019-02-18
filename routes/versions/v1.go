package versions

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/http/middleware"
	"github.com/totoval/framework/route"
	"totoval/routes/groups"
)

type V1 struct {
	Prefix string
}

func (v1 *V1) Register(router *gin.Engine) {
	version := router.Group(v1.Prefix)
	{
		noAuth(version)
		auth(version)
	}
}

func noAuth(group *gin.RouterGroup) {
	route.RegisterRouteGroup(&groups.AuthGroup{}, group)
}

func auth(group *gin.RouterGroup) {
	authGroup := group.Group("", middleware.AuthRequired())

	{
		route.RegisterRouteGroup(&groups.UserGroup{}, authGroup)
	}
}

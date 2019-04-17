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
		v1.noAuth(version)
		v1.auth(version)
	}
}

func (v1 *V1) noAuth(group *gin.RouterGroup) {
	noAuthGroup := group.Group("")

	{
		route.RegisterRouteGroup(&groups.AuthGroup{}, noAuthGroup)
		route.RegisterRouteGroup(&groups.UserAffiliationGroup{}, noAuthGroup)
	}
}

func (v1 *V1) auth(group *gin.RouterGroup) {
	authGroup := group.Group("", middleware.AuthRequired())

	{
		route.RegisterRouteGroup(&groups.UserGroup{}, authGroup)
	}
}

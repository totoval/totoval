package groups

import (
	"github.com/gin-gonic/gin"

	"totoval/app/http/controllers"
)

type UserAffiliationGroup struct {
	UserAffiliationController controllers.UserAffiliation
}

func (uaffg *UserAffiliationGroup) Register(group *gin.RouterGroup) {

	newGroup := group.Group("/user-affiliation")
	{
		newGroup.GET("/all", uaffg.UserAffiliationController.RenderAll)
	}
}

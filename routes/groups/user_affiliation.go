package groups

import (
	"github.com/totoval/framework/route"
	"totoval/app/http/controllers"
)

type UserAffiliationGroup struct {
	UserAffiliationController controllers.UserAffiliation
}

func (uaffg *UserAffiliationGroup) Group(group route.Grouper) {
	group.GET("/all", uaffg.UserAffiliationController.RenderAll)
}

package versions

import (
	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/route"
	"totoval/routes/groups"
)

func NewV1(engine *gin.Engine) {
	ver := route.NewVersion(engine, "v1")

	// auth routes
	ver.Auth("", func(grp route.Grouper) {
		grp.AddGroup("/user", &groups.UserGroup{})
	})

	// no auth routes
	ver.NoAuth("", func(grp route.Grouper) {
		grp.AddGroup("", &groups.AuthGroup{})
		grp.AddGroup("/user-affiliation", &groups.UserAffiliationGroup{})
	})
}

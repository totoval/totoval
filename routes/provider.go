package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/route"
	"totoval/routes/versions"
)

func Register(router *gin.Engine) {
	defer route.Bind()

	versions.NewV1(router)
}

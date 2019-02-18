package routes

import (
	"github.com/gin-gonic/gin"
	"totoval/routes/versions"
)

func Register(router *gin.Engine) {
	version := &versions.V1{Prefix: "v1"}

	version.Register(router)
}

package routes

import (
	"github.com/gin-gonic/gin"

	"totoval/routes/versions"
)

func Register(router *gin.Engine) {
	versions.NewV1(router)
}

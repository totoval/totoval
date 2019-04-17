package views

import (
	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/view"
)

func Initialize(r *gin.Engine) {
	view.Initialize(r)
}

package config

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/config"
)

func Initialize() {
	setAppMode()
}

func setAppMode() {
	if config.GetString("app.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

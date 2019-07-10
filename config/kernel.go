package config

import (
	"github.com/totoval/framework/app"
	"github.com/totoval/framework/config"
)

func Initialize() {
	setAppMode()
}

func setAppMode() {
	switch config.GetString("app.env") {
	case "production":
		app.SetMode(app.ModeProduction)
	case "develop":
		app.SetMode(app.ModeDevelop)
	case "test":
		app.SetMode(app.ModeTest)
	default:
		app.SetMode(app.ModeDevelop)
	}
}

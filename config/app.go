package config

import (
	. "totoval-framework/config"
)

func init() {
	app := make(map[string]interface{})

	app["name"] = Env("APP_NAME", "Totoval")
	app["env"] = Env("APP_ENV", "production")
	app["debug"] = Env("APP_DEBUG", false)
	app["timezone"] = "Asia/Shanghai"
	app["locale"] = Env("APP_LOCALE", "en")
	app["fallback_locale"] = "en"
	app["key"] = Env("APP_KEY")
	app["cipher"] = "AES-256-CBC"

	Add("app", app)
}

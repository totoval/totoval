package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "sign key")

	Add("auth", auth)
}

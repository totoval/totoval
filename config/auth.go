package config

import (
	. "github.com/totoval/framework/config"
	"totoval/app/models"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "sign key")
	auth["model_ptr"] = &models.User{} // must be a pointer

	Add("auth", auth)
}

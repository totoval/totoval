package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	userAffiliation := make(map[string]interface{})

	userAffiliation["enable"] = true

	Add("user_affiliation", userAffiliation)
}

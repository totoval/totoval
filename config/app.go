package config

import (
	"totoval-framework/config"
)

func init(){
	conf := make(map[string]interface{})
	conf["test"] = 123
	conf["test1"] = config.Env("test1", "test1 not set default")
	conf["test2"] = []string{
		"123", "456",
	}
	config.Add("app", conf)

}


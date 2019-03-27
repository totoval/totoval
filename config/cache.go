package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	cache := make(map[string]interface{})

	cache["default"] = Env("CACHE_DRIVER", "memory")
	cache["stores"] = map[string]interface{}{
		"memory": map[string]interface{}{
			"driver":                    "memory",
			"default_expiration_minute": 5,
			"cleanup_interval_minute":   5,
			"prefix":                    Env("CACHE_PREFIX", "totoval_"),
		},
	}

	Add("cache", cache)
}

package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	throttle := make(map[string]interface{})

	throttle["max_attempts"] = 60
	throttle["decay_minutes"] = 1

	Add("throttle", throttle)
}

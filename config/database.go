package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	database := make(map[string]interface{})

	database["default"] = Env("DB_CONNECTION", "mysql")
	database["connections"] = map[string]interface{}{
		"mysql": map[string]interface{}{
			"driver":    "mysql",
			"host":      Env("DB_HOST", "127.0.0.1"),
			"port":      Env("DB_PORT", "3306"),
			"database":  Env("DB_DATABASE", ""),
			"username":  Env("DB_USERNAME", ""),
			"password":  Env("DB_PASSWORD", ""),
			"charset":   "utf8mb4",
			"collation": "utf8mb4_unicode_ci",
			"prefix":    Env("DB_PREFIX", ""),
		},
	}
	database["migrations"] = "migrations"
	database["max_idle_connections"] = 10
	database["max_open_connections"] = 100

	database["redis"] = map[string]interface{}{
		"default": map[string]interface{}{
			"host":     Env("REDIS_HOST", "127.0.0.1"),
			"password": Env("REDIS_PASSWORD", nil),
			"port":     Env("REDIS_PORT", 6379),
			"database": Env("REDIS_DB", 0),
		},
		"cache": map[string]interface{}{
			"host":     Env("REDIS_HOST", "127.0.0.1"),
			"password": Env("REDIS_PASSWORD", nil),
			"port":     Env("REDIS_PORT", 6379),
			"database": Env("REDIS_CACHE_DB", 1),
		},
	}

	Add("database", database)
}

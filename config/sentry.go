package config

import (
	. "github.com/totoval/framework/config"
)

func init() {
	sentry := make(map[string]interface{})

	sentry["enable"] = Env("SENTRY_ENABLE", false)
	sentry["host"] = Env("SENTRY_HOST", "app.getsentry.com")
	sentry["key"] = Env("SENTRY_KEY", "YOUR-OWN-SENTRY-KEY")
	sentry["secret"] = Env("SENTRY_SECRET", "YOUR-OWN-SENTRY-SECRET")
	sentry["project"] = Env("SENTRY_PROJECT", "YOUR-OWN-SENTRY-PROJECT")

	Add("sentry", sentry)
}

package bootstrap

import (
	"github.com/totoval/framework/cache"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/helpers/zone"
	"github.com/totoval/framework/logs"
	"github.com/totoval/framework/queue"
	"github.com/totoval/framework/utils/validator"
	"totoval/app/events"
	"totoval/app/jobs"
	"totoval/app/listeners"
	"totoval/config"
	"totoval/resources/lang"
)

func Initialize() {
	config.Initialize()
	logs.Initialize()
	zone.Initialize()
	cache.Initialize()
	database.Initialize()
	m.Initialize()
	lang.Initialize() // an translation must contains resources/lang/xx.json file (then a resources/lang/validation_translator/xx.go)
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()
	
	validator.UpgradeValidatorV8toV9()
}

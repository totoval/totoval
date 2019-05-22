package main

import (
	"context"
	"os"
	"time"

	"github.com/totoval/framework/graceful"
	"github.com/totoval/framework/helpers/log"

	"github.com/urfave/cli"

	"github.com/totoval/framework/console"
	"github.com/totoval/framework/helpers/zone"
	"github.com/totoval/framework/logs"

	"totoval/database/migrations"

	"github.com/totoval/framework/cache"
	"github.com/totoval/framework/cmd"
	command_queue "github.com/totoval/framework/cmd/commands/queue"
	"github.com/totoval/framework/cmd/commands/schedule"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/queue"

	app_schedule "totoval/app/console"

	"totoval/app/console/commands"

	"totoval/app/events"
	"totoval/app/jobs"
	"totoval/app/listeners"
	"totoval/config"
	"totoval/resources/lang"
)

func init() {
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

	migrations.Initialize()
	command_queue.Initialize()
	commands.Initialize()
	schedule.Initialize()
}

func main() {
	scheduleInit()

	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "Let's work like an artisan"
	app.Version = "0.5.5"

	app.Commands = cmd.List()

	app.Action = func(c *cli.Context) error {
		console.Println(console.CODE_INFO, "COMMANDS:")

		for _, cate := range app.Categories() {
			categoryName := cate.Name
			if categoryName == "" {
				categoryName = "kernel"
			}
			console.Println(console.CODE_WARNING, "    "+categoryName+":")

			for _, cmds := range cate.Commands {
				console.Println(console.CODE_SUCCESS, "        "+cmds.Name+" "+console.Sprintf(console.CODE_INFO, "%s", cmds.ArgsUsage)+"    "+console.Sprintf(console.CODE_WARNING, "%s", cmds.Usage))
			}
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer graceful.ShutDown(ctx)
	defer cancel()
}

func scheduleInit() {
	app_schedule.Schedule(cmd.NewSchedule())
}

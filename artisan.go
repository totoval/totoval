package main

import (
	"os"

	"github.com/totoval/framework/graceful"
	"github.com/totoval/framework/helpers/log"
	"github.com/totoval/framework/sentry"

	"totoval/bootstrap"

	"github.com/urfave/cli"

	"github.com/totoval/framework/console"

	"totoval/database/migrations"

	"github.com/totoval/framework/cmd"
	command_queue "github.com/totoval/framework/cmd/commands/queue"
	"github.com/totoval/framework/cmd/commands/schedule"

	app_schedule "totoval/app/console"

	"totoval/app/console/commands"
)

func init() {
	bootstrap.Initialize()

	migrations.Initialize()
	command_queue.Initialize()
	commands.Initialize()
	schedule.Initialize()

	app_schedule.Schedule(cmd.NewSchedule())
}

func main() {
	cliServe()
}

func cliServe() {
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
		sentry.CaptureError(err)
		log.Fatal(err.Error())
	}

	// totoval framework shutdown
	graceful.ShutDown(true)
}

package main

import (
	"github.com/totoval/framework/cmd"
	"github.com/totoval/framework/cmd/groups"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/urfave/cli"
	"log"
	"os"
	"totoval/config"
	"totoval/database/migrations"
)

func init() {
	config.Initialize()
	database.Initialize()
	m.Initialize()
}

func main() {
	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "Let's work like an artisan"

	chLog := make(chan interface{})

	// command group
	migrateCommand := &groups.MigrateCommand{MigratorInitializer: migrations.Initialize, ChLog: chLog}

	//app.Flags = []cli.Flag {
	//	cli.StringFlag{
	//		Name:        "lang",
	//		Value:       "english",
	//		Usage:       "language for the greeting",
	//		Destination: &language,
	//	},
	//}

	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Commands = []cli.Command{
		migrateCommand.MigrationInit(),
		migrateCommand.Migrate(),
		migrateCommand.MigrateRollBack(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	receiveLog(chLog)
}

func receiveLog(chLog chan interface{}) {
	for _log := range chLog {
		if _log == nil {
			os.Exit(0)
		}
		if __log, ok := _log.(cmd.TermLog); ok {
			__log.Print()
		}
	}
}

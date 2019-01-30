package main

import (
	"Wallet/database/migrations"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main(){
	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "Let's work like an artisan"

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
		{
			Name:    "migrate",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action:  func(c *cli.Context) error {
				m := &migration.MigrationUtils{}
				m.SetDB()
				//m.Initialize()
				m.Migrate()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
package cmd

import (
	"github.com/mingrammer/cfmt"
	"github.com/urfave/cli/v2"
	"os"
)

func Run() {

	app := &cli.App{
		Name:    "gyv",
		Usage:   "The web framework with a special flavour.",
		Version: "v0.1.0",
		Authors: []*cli.Author{
			{
				Name:  "Alexander Wulf",
				Email: "hi@alexfwulf.de",
			},
		},
		Commands: []*cli.Command{
			initCmd,
			makeControllerCmd,
			makeMiddlewareCmd,
			makeModelCmd,
			makeRequestCmd,
			migrateFreshCmd,
			migrateRunCmd,
			routesCmd,
			seedCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = cfmt.Errorln("Error: ", err)
	}
}

package cmd

import (
	"os"

	"github.com/mingrammer/cfmt"
	"github.com/urfave/cli/v2"
)

func Run() {

	app := &cli.App{
		Name:    "gyv",
		Usage:   "The goyave-cli with a special flavour.",
		Version: "v0.1.0-nightly",
		Authors: []*cli.Author{
			{
				Name:  "Alexander Wulf",
				Email: "dev@alexfwulf.de",
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

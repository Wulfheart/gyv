package cmd

import "github.com/urfave/cli/v2"

var migrateRunCmd = &cli.Command{
	Name:     "migrate:run",
	Category: "migrate",
	Action: func(c *cli.Context) error {
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

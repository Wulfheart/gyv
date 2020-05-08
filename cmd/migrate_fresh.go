package cmd

import "github.com/urfave/cli/v2"

var migrateFreshCmd = &cli.Command{
	Name:     "migrate:fresh",
	Category: "migrate",
	Action: func(c *cli.Context) error {
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

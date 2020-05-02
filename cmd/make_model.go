package cmd

import "github.com/urfave/cli/v2"

var makeModelCmd = &cli.Command{
	Name:      "make:model",
	Category:  "make",
	Action:   func(c *cli.Context) error {
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

package cmd

import "github.com/urfave/cli/v2"

var makeRequestCmd = &cli.Command{
	Name:      "make:request",
	Category:  "make",
	Action:    func(c *cli.Context) error {
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

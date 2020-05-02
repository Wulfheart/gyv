package cmd

import "github.com/urfave/cli/v2"

var makeMiddlewareCmd = &cli.Command{
	Name:     "make:middleware",
	Category: "make",
	Action: func(c *cli.Context) error {
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

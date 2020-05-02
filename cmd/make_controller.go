package cmd

import (
	"github.com/urfave/cli/v2"
)

var makeControllerCmd = &cli.Command{
	Name:      "make:controller",
	Category:  "make",
	Action: func(context *cli.Context) error {
		// err := changeWorkingDirectoryToRoot()
		// if err != nil {
		// 	return err
		// }
		notImplemented()
		return nil
	},
	Usage:     "WIP",
	UsageText: "WIP",
}

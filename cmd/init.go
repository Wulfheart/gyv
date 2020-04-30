package cmd

import "github.com/urfave/cli/v2"

var initCmd = &cli.Command{
	Name:      "init",
	Action:    nil,
	Usage:     "Initializes a new project",
	UsageText: "Text is a little bit longer",
}

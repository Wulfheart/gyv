package cmd

import (
	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name: "init",
	Action: func(c *cli.Context) error {
		//  check if directory name already exists
		// Download template project (https://github.com/System-Glitch/goyave-template/archive/master.zip)
		// unzip template
		// move content to folder
		// copy config.example.json to config.json
		// Initialize git

		//! We need another template repo as I don't want to use the find function

		//
		return nil
	},
	Usage:     "Initializes a new project",
	UsageText: "Text is a little bit longer",
}

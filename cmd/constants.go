package cmd

import "github.com/mingrammer/cfmt"

const (
	placeholderUsage     = "This is a placeholder"
	placeholderUsageText = "This is a longer text"
)

func notImplemented() {
	_, _ = cfmt.Infoln("This command is currently not implemented.\n" +
		"Please check current status at https://github.com/Wulfheart/gyv.")
}

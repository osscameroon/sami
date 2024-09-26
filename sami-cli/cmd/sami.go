package cmd

import (
	"github.com/spf13/cobra"
)

//NewCommand returns a new sami cobra command
func NewCommand() *cobra.Command {
	var samiCommand = &cobra.Command{
		Use:   "sami",
		Short: "sami helps you deploy your services to any target machine",
		Long:  `sami helps you deploy your services to any target machine`,
	}

	// deploy
	// samiCommand.AddCommand(NewDeployCommand())

	// logs
	// samiCommand.AddCommand(NewLogsCommand())

	// oob
	// samiCommand.AddCommand(NewOobCommand())

	// status
	// samiCommand.AddCommand(NewStatusCommand())
	return samiCommand
}

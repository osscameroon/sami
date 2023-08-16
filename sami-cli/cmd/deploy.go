package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

//NewDeployCommand create new cobra command for the init command
func NewDeployCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "deploy -f <deployment_sami_file>",
		Short: "deloy your service by given the deployment sami file",
		Long: `deloy your service by given the deployment sami file

		example: sami deploy -f <deploy-path>/<sami.yaml>
`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deployFolder := args[0]

			if err := deployCommand(deployFolder); err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
		},
	}

	return command
}

func deployCommand(deployFolder string) error {
	// will add the deployment process here...
	return errors.New("")
}

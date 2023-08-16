package main

import (
	"os"

	"github.com/osscameroon/sami/sami-cli/cmd"
)

var version = "0.0.0+dev"

func main() {
	command := cmd.NewCommand()

	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

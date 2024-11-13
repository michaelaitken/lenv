package main

import (
	"os"

	"github.com/michaelaitken/lenv/pkg/env"
	"github.com/michaelaitken/lenv/pkg/profile"
	"github.com/michaelaitken/lenv/utils"
)

func main() {
	// Seperate and manage the CLI argments
	if len(os.Args) < 2 {
		utils.PrintUsageInformation()
		os.Exit(1)
	}

	// appPath := os.Args[0]
	command := os.Args[1]
	commandVariables := os.Args[2:]

	switch command {
	// Initialise the local environment directory file
	case "init":
		env.InitCmd()

	// Create a new environment profile with the specified name
	case "create":
		if utils.CheckCommandVariables(1, &commandVariables) {
			profile.CreateCmd(commandVariables)
		}

	// Activate an existing environment with the specified name
	case "activate":
		if utils.CheckCommandVariables(1, &commandVariables) {
			profile.ActivateCmd(commandVariables)
		}

	// Deactivates the active environment profile
	case "deactivate":
		profile.DeactivateCmd()

	// Print help information to the terminal
	case "help":
		utils.PrintHelpInformation()

	// Print usage information when detecting unrecognized command
	default:
		utils.PrintUsageInformation()
	}
}

package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/michaelaitken/lenv/pkg/parser"
	"github.com/michaelaitken/lenv/pkg/script"
	"github.com/michaelaitken/lenv/utils"
)

func ActivateCmd(variables []string) {
	path := utils.GetEnvDirectoryPath()
	profileName := variables[0]
	currentProfile, exists := os.LookupEnv("LENV_ENVIRONMENT_NAME")

	// Check for active environment
	if exists && currentProfile == profileName {
		fmt.Fprintf(os.Stderr, "%s profile is already active!", currentProfile)
		os.Exit(0)
	} else if exists {
		fmt.Fprintf(os.Stderr, "Please deactivate the active '%s' environment!", currentProfile)
		os.Exit(0)
	}

	// Check changes to the YAML file
	scriptFilePath := filepath.Join(path, "scripts", fmt.Sprintf("script-%s.ps1", profileName))
	scriptFileInfo, err := os.Stat(scriptFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find script file: %v", err)
		os.Exit(1)
	}

	profileFilePath := filepath.Join(path, fmt.Sprintf("%s.yaml", profileName))
	profileFileInfo, err := os.Stat(profileFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find profile yaml file: %v", err)
		os.Exit(1)
	}

	// If not changes, just run the current script file
	if isTrue := profileFileInfo.ModTime().After(scriptFileInfo.ModTime()); !isTrue {
		script.Execute(path, profileName)
		os.Exit(0)
	}

	// Regenerate script file if changes are detected.
	fmt.Println("Changes to YAML file detected!")

	file, err := os.Open(profileFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open profile yaml file: %v", err)
		os.Exit(1)
	}

	varMap, err := parser.ParseYamlFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse profile yaml file: %v", err)
		os.Exit(1)
	}

	err = script.Generate(varMap, profileName, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot generate new script file: %v", err)
		os.Exit(1)
	}

	// Execute the newly generated script
	script.Execute(path, profileName)
	os.Exit(0)
}

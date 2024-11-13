package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/michaelaitken/lenv/pkg/script"
	"github.com/michaelaitken/lenv/utils"
)

func CreateCmd(variables []string) {
	path := utils.GetEnvDirectoryPath()
	profileName := variables[0]

	// Create profile YAML file
	profile := profileName + ".yaml"

	profilePath := filepath.Join(path, profile)
	_, err := os.Create(profilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create profile file: %v", err)
		os.Exit(1)
	}

	err = script.Generate(make(map[string]string), profileName, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create profile script: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

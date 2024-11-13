package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/michaelaitken/lenv/utils"
)

func InitCmd() {
	// Get Working Directory
	envDir := utils.GetEnvDirectoryPath()

	// Create the directory if it doesn't exist
	_, err := os.Stat(envDir)
	if err == nil {
		fmt.Println("Local Environment already initialized.")
		os.Exit(0)
	} else if !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "unknown error while checking .lenv directory: %v\n", err)
		os.Exit(1)
	}

	// Define a readable file mode for directory creation (rwxr-xr-x)
	const dirPerm = os.ModeDir | 0755

	// Create the .lenv directory
	if err := os.Mkdir(envDir, dirPerm); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create .lenv directory: %v\n", err)
		os.Exit(1)
	}

	// Create the scripts subdirectory
	envScriptsDir := filepath.Join(envDir, "scripts")
	if err := os.Mkdir(envScriptsDir, dirPerm); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create scripts directory: %v\n", err)
		os.Exit(1)
	}
}

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetEnvDirectoryPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get current directory: %v\n", err)
		os.Exit(1)
	}
	envDir := filepath.Join(currentDir, ".lenv")
	return envDir
}

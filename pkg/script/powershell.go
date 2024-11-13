package script

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Generate creates a PowerShell script to toggle environment variables.
func Generate(env map[string]string, envName string, envDir string) error {
	// Add Environment Identifier to env
	env[EnvVariableNameIdentifier] = envName

	// Define the script path and create/overwrite the file
	scriptPath := filepath.Join(envDir, fmt.Sprintf("%s-%s.ps1", ScriptPath, envName))
	file, err := os.Create(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to create PowerShell script: %w", err)
	}
	defer file.Close()

	// Write the PowerShell toggle logic
	file.WriteString("$isSet = $false\n")
	for key := range env {
		file.WriteString(fmt.Sprintf("if ($env:%s -ne $null) { $isSet = $true }\n", key))
	}

	// Logic to enable or disable environment variables
	file.WriteString("if ($isSet) {\n")
	for key := range env {
		file.WriteString(fmt.Sprintf("    Remove-Item Env:\\%s -ErrorAction SilentlyContinue\n", key))
	}
	file.WriteString("    exit\n} else {\n")
	for key, value := range env {
		file.WriteString(fmt.Sprintf("    $env:%s = \"%s\"\n", key, value))
	}
	file.WriteString("    Write-Host 'Local environment variables enabled.' -ForegroundColor Green\n}\n")

	// Define a custom prompt function to reflect environment state
	file.WriteString("function localPrompt {\n ")
	file.WriteString("	$breadcrumb = if (\n    ")
	last := len(env)
	count := 1
	for key := range env {
		file.WriteString(fmt.Sprintf("$env:%s", key))
		if count != last {
			file.WriteString(" -and \n")
		}
		count++
	}
	file.WriteString(") { \"$([char]27)[0;92m(.lenv-" + envName + ")$([char]27)[0m \" } else { \"\" }\n")
	file.WriteString("    return \"$breadcrumb$(Get-Location)> \"\n}\n")

	// Set the PowerShell prompt and clean up
	file.WriteString("Set-Item -Path Function:\\global:prompt -Value (Get-Command localPrompt).Definition\n")
	file.WriteString("Remove-Item -Path Function:\\localPrompt -ErrorAction SilentlyContinue\n")

	return nil
}

// Execute the PowerShell script to activate the environment
func Execute(envDir string, profileName string) error {
	path := filepath.Join(envDir, fmt.Sprintf("%s-%s.ps1", ScriptPath, profileName))

	cmd := exec.Command("powershell", "-NoExit", "-Command", fmt.Sprintf("& { %s; }", path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Used to make it interactive

	return cmd.Run()
}

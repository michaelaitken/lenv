package utils

import "fmt"

func PrintUsageInformation() {
	fmt.Println("Use 'lenv help' for usage information.")
}

func PrintHelpInformation() {
	fmt.Println("lenv help information.")
}

func CheckCommandVariables(numRequiredVars int, variableSlice *[]string) bool {
	return len(*variableSlice) == numRequiredVars
}

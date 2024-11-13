package parser

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ParseYamlFile reads a simple YAML file and returns a map of variables
func ParseYamlFile(file *os.File) (map[string]string, error) {
	varMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse key-value pairs (key: value)
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid YAML syntax: missing ':' delimiter")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes from the value, if any
		value = removeDoubleQuotes(value)

		varMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return varMap, nil
}

// Helper function to remove double quotes from a string, if present
func removeDoubleQuotes(s string) string {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return s[1 : len(s)-1]
	}
	return s
}

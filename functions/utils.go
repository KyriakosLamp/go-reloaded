package functions

import (
	"os"
)

// ReadFile reads the entire content of a file and returns it as a string
func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes string content to a file with standard permissions (0644)
func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
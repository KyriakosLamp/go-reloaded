package functions

import (
	"fmt"
	"log"
	"os"
)

// Logger handles controlled warnings and error logging
type Logger struct {
	enabled bool
}

// NewLogger creates a new logger instance
func NewLogger(enabled bool) *Logger {
	return &Logger{enabled: enabled}
}

// LogWarning logs a warning message if logging is enabled
func (l *Logger) LogWarning(message string) {
	if l.enabled {
		log.Printf("WARNING: %s", message)
	}
}

// LogError logs an error message and exits if critical
func (l *Logger) LogError(message string, critical bool) {
	if critical {
		fmt.Printf("ERROR: %s\n", message)
		os.Exit(1)
	} else if l.enabled {
		log.Printf("ERROR: %s", message)
	}
}

// Global logger instance
var GlobalLogger = NewLogger(false) // Disabled by default for clean output
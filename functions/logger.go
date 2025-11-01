package functions

import (
	"fmt"
	"log"
	"os"
)

// Logger provides configurable warning and error logging functionality
type Logger struct {
	enabled bool
}

// NewLogger creates a new logger instance with specified enable state
func NewLogger(enabled bool) *Logger {
	return &Logger{enabled: enabled}
}

// LogWarning logs a warning message if logging is enabled
func (l *Logger) LogWarning(message string) {
	if l.enabled {
		log.Printf("WARNING: %s", message)
	}
}

// LogError logs an error message and optionally exits if critical
func (l *Logger) LogError(message string, critical bool) {
	if critical {
		fmt.Printf("ERROR: %s\n", message)
		os.Exit(1)
	} else if l.enabled {
		log.Printf("ERROR: %s", message)
	}
}

// GlobalLogger is the shared logger instance, disabled by default for clean CLI output
var GlobalLogger = NewLogger(false)
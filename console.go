package persisters

import (
	"fmt"

	logging "github.com/remoteit/systemkit-logging"
)

type consoleLogger struct{}

// NewConsoleLogger -
func NewConsoleLogger() logging.CoreLogger {
	return &consoleLogger{}
}

func (thisRef consoleLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	fmt.Println(logEntry.Message)

	return logEntry
}

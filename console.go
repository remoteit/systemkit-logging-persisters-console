package persisters

import (
	"log"

	logging "github.com/codemodify/systemkit-logging"
	"github.com/mattn/go-colorable"
)

type consoleLogger struct{}

// NewConsoleLogger -
func NewConsoleLogger() logging.CoreLogger {
	return &consoleLogger{}
}

func (thisRef consoleLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	log.SetOutput(colorable.NewColorableStdout()) // or NewColorableStderr()
	log.SetFlags(0)

	log.Println(logEntry.Message)

	return logEntry
}

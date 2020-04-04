package persisters

import (
	"log"

	logging "github.com/codemodify/systemkit-logging"
	"github.com/mattn/go-colorable"
)

type consoleLogger struct {
	colors map[logging.LogType]ConsoleLoggerColorDelegate
}

// NewConsoleLogger -
func NewConsoleLogger(colors map[logging.LogType]ConsoleLoggerColorDelegate) logging.Logger {
	return &consoleLogger{
		colors: colors,
	}
}

func (thisRef consoleLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	log.SetOutput(colorable.NewColorableStdout()) // or NewColorableStderr()
	log.SetFlags(0)

	if logEntry.Type == logging.TypeTrace {
		log.Println(thisRef.colors[logging.TypeTrace](logEntry.Message))

	} else if logEntry.Type < logging.TypeWarning {
		log.Println(thisRef.colors[logging.TypeError](logEntry.Message))

	} else if logEntry.Type == logging.TypeWarning {
		log.Println(thisRef.colors[logging.TypeWarning](logEntry.Message))

	} else if logEntry.Type == logging.TypeInfo {
		log.Println(thisRef.colors[logging.TypeInfo](logEntry.Message))

	} else if logEntry.Type == logging.TypeSuccess {
		log.Println(thisRef.colors[logging.TypeSuccess](logEntry.Message))

	} else if logEntry.Type == logging.TypeDebug {
		log.Println(thisRef.colors[logging.TypeDebug](logEntry.Message))
	}

	return logEntry
}

package persisters

import (
	"log"

	logging "github.com/codemodify/systemkit-logging"
	"github.com/mattn/go-colorable"
)

type consoleLogger struct {
	logUntil logging.LogType
	colors   map[logging.LogType]ConsoleLoggerColorDelegate
}

// NewConsoleLogger -
func NewConsoleLogger(logUntil logging.LogType, colors map[logging.LogType]ConsoleLoggerColorDelegate) logging.Logger {
	return &consoleLogger{
		logUntil: logUntil,
		colors:   colors,
	}
}

// NewConsoleLoggerDefaultColors -
func NewConsoleLoggerDefaultColors() map[logging.LogType]ConsoleLoggerColorDelegate {
	return map[logging.LogType]ConsoleLoggerColorDelegate{
		logging.TypeDisable: WhiteString,
		logging.TypeTrace:   BlackStringYellowBG,
		logging.TypePanic:   RedString,
		logging.TypeFatal:   RedString,
		logging.TypeError:   RedString,
		logging.TypeWarning: YellowString,
		logging.TypeInfo:    WhiteString,
		logging.TypeSuccess: GreenString,
		logging.TypeDebug:   CyanString,
	}
}

// ConsoleLoggerColorDelegate -
type ConsoleLoggerColorDelegate func(string, ...interface{}) string

// BlackStringYellowBG -
func BlackStringYellowBG(format string, a ...interface{}) string {
	c := New(FgBlack, BgYellow)
	return c.Sprintf(format, a...)
}

// BlackStringWhiteBG -
func BlackStringWhiteBG(format string, a ...interface{}) string {
	c := New(FgBlack, BgWhite)
	return c.Sprintf(format, a...)
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

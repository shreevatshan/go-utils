package log

import (
	"fmt"
	"sync"
	"time"
)

const (
	defaultBufferSize = 1024
)

/*
log level
*/
const (
	NoLog = iota
	Info
	Warning
	Debug
	Trace
)

/*
log message prefix
*/
const (
	msgPrefixInfo    = "INFO"
	msgPrefixWarning = "WARNING"
	msgPrefixDebug   = "DEBUG"
	msgPrefixTrace   = "DEV"
	msgPrefixPanic   = "PANIC"
)

type Logger interface {
	InfoLog(format string, args ...interface{})
	DebugLog(format string, args ...interface{})
	WarningLog(format string, args ...interface{})
	TraceLog(format string, args ...interface{})
	PanicQuickLog(format string, args ...interface{})
	QuickLog(level int, format string, args ...interface{})
	LogMessage(level int, format string, args ...interface{})
}

var (
	loggerCollection map[string]Logger
	mu               sync.RWMutex
)

func init() {
	loggerCollection = make(map[string]Logger)
}

func RegisterLogger(loggerName string, logger Logger) {
	mu.Lock()
	loggerCollection[loggerName] = logger
	mu.Unlock()
}

func GetLogger(loggerName string) Logger {
	mu.RLock()
	defer mu.RUnlock()

	if _, exists := loggerCollection[loggerName]; exists {
		return loggerCollection[loggerName]
	}
	return &StdLogger{}
}

func formatMessage(level string, msg string) string {
	now := time.Now()
	return fmt.Sprintf("[%s]\t[%s]\t%s\n", level, now.Format(time.ANSIC), msg)
}

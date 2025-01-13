package log

import (
	"fmt"
	"sync"
	"time"
)

const (
	defaultBuffer   = 1024
	defaultSize     = 5000000
	defaultMaxFiles = 10
)

type Logger interface {
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Panic(format string, args ...interface{})
	Quick(level Level, format string, args ...interface{})
	WithLevel(level Level, format string, args ...interface{})
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
	defer mu.Unlock()
	loggerCollection[loggerName] = logger
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

// log level
const (
	NoLog = iota
	Info
	Warning
	Debug
	Trace
)

const (
	nolog   = "NOLOG"
	info    = "INFO"
	warning = "WARNING"
	debug   = "DEBUG"
	trace   = "TRACE"
	panic   = "PANIC"
)

type Level int

func (l Level) String() string {
	switch l {
	case Info:
		return info
	case Warning:
		return warning
	case Debug:
		return debug
	case Trace:
		return trace
	default:
		return "UNKNOWN"
	}
}

func ToLevel(level string) Level {
	switch level {
	case nolog:
		return NoLog
	case info:
		return Info
	case warning:
		return Warning
	case debug:
		return Debug
	case trace:
		return Trace
	default:
		return Warning
	}
}

package log

import "fmt"

type StdLogger struct{}

func (l *StdLogger) Info(format string, args ...interface{}) {
	fmt.Println(formatMessage(info, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	fmt.Println(formatMessage(debug, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) Warning(format string, args ...interface{}) {
	fmt.Println(formatMessage(warning, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) Trace(format string, args ...interface{}) {
	fmt.Println(formatMessage(trace, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) Panic(format string, args ...interface{}) {
	fmt.Println(formatMessage(panic, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) Quick(level Level, format string, args ...interface{}) {
	fmt.Println(formatMessage(level.String(), fmt.Sprintf(format, args...)))
}

func (l *StdLogger) WithLevel(level Level, format string, args ...interface{}) {
	fmt.Println(formatMessage(level.String(), fmt.Sprintf(format, args...)))
}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

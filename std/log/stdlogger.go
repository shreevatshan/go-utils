package log

import "fmt"

type StdLogger struct{}

func (l *StdLogger) InfoLog(format string, args ...interface{}) {
	fmt.Println(formatMessage(msgPrefixInfo, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) DebugLog(format string, args ...interface{}) {
	fmt.Println(formatMessage(msgPrefixDebug, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) WarningLog(format string, args ...interface{}) {
	fmt.Println(formatMessage(msgPrefixWarning, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) TraceLog(format string, args ...interface{}) {
	fmt.Println(formatMessage(msgPrefixTrace, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) PanicQuickLog(format string, args ...interface{}) {
	fmt.Println(formatMessage(msgPrefixPanic, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) QuickLog(level int, format string, args ...interface{}) {
	var prefix string
	switch level {
	case Info:
		prefix = msgPrefixInfo
	case Warning:
		prefix = msgPrefixWarning
	case Debug:
		prefix = msgPrefixDebug
	case Trace:
		prefix = msgPrefixTrace
	default:
		return
	}
	fmt.Println(formatMessage(prefix, fmt.Sprintf(format, args...)))
}

func (l *StdLogger) LogMessage(level int, format string, args ...interface{}) {
	var prefix string
	switch level {
	case Info:
		prefix = msgPrefixInfo
	case Warning:
		prefix = msgPrefixWarning
	case Debug:
		prefix = msgPrefixDebug
	case Trace:
		prefix = msgPrefixTrace
	default:
		return
	}
	fmt.Println(formatMessage(prefix, fmt.Sprintf(format, args...)))
}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

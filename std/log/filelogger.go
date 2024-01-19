package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type FileLog struct {
	receive           bool
	loglevel          int
	logsize           int
	buffersize        int
	filename          string
	location          string
	done              chan struct{}
	rotate            chan struct{}
	buffer            chan string
	logfiledescriptor *os.File
	wg                sync.WaitGroup
	mu                sync.Mutex
	currentsize       atomic.Int32
}

/*
syntax:

	Init(logFilename string, logLocation string, logLevel int, logSize int, bufferSize int)

mandatory:

	logFilename, logLocation, logLevel

optional:

	logSize, bufferSize
*/
func InitFileLogger(logFilename string, logLocation string, logLevel int, args ...interface{}) *FileLog {

	var logSize int
	var bufferSize = defaultBufferSize

	for index, arg := range args {
		switch index {
		case 0:
			if arg != nil {
				logSize = arg.(int)
			}
		case 1:
			if arg != nil {
				bufferSize = arg.(int)
			}
		}
	}

	logger := &FileLog{
		loglevel:   logLevel,
		filename:   logFilename,
		location:   logLocation,
		logsize:    logSize,
		buffersize: bufferSize,
		buffer:     make(chan string, bufferSize),
		done:       make(chan struct{}, 1),
		rotate:     make(chan struct{}, 1),
	}

	return logger
}

/*
syntax:

	Update(logFilename string, logLocation string, logLevel int, logSize int, bufferSize int)

mandatory:

	logFilename, logLocation, logLevel

optional:

	logSize, bufferSize
*/
func (logger *FileLog) Update(logFilename string, logLocation string, logLevel int, args ...interface{}) error {
	var err error
	var bufferSize = defaultBufferSize

	logSize := logger.logsize

	for index, arg := range args {
		switch index {
		case 0:
			if arg != nil {
				logSize = arg.(int)
			}
		case 1:
			if arg != nil {
				bufferSize = arg.(int)
			}
		}
	}

	if (logger.filename != logFilename) || (logger.location != logLocation) || (logger.loglevel != logLevel) || (logger.logsize != logSize) || (logger.buffersize != bufferSize) {
		logger.Stop()

		logger.filename = logFilename
		logger.location = logLocation
		logger.loglevel = logLevel
		logger.logsize = logSize
		logger.buffersize = bufferSize
		logger.buffer = make(chan string, bufferSize)

		err = logger.Start()
	}
	return err
}

func (logger *FileLog) open() error {
	var err error

	err = os.MkdirAll(logger.location, 0777)
	if err != nil {
		return err
	}

	logfilename := filepath.Join(logger.location, logger.filename)

	logger.logfiledescriptor, err = os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return err
}

func (logger *FileLog) close() error {
	return logger.logfiledescriptor.Close()
}

func (logger *FileLog) Start() error {

	logger.mu.Lock()

	if logger.receive {
		return fmt.Errorf("logger already in running state")
	}

	err := logger.open()

	if err == nil {
		logger.logfiledescriptor.Write([]byte(formatMessage(msgPrefixInfo, "STARTED")))
		logger.receive = true
		logger.wg.Add(1)
		go func() {
			logger.flusher()
			logger.wg.Done()
		}()
	}

	logger.mu.Unlock()

	return err
}

func (logger *FileLog) Stop() {

	logger.mu.Lock()

	if logger.receive {
		logger.receive = false

		if len(logger.done) < 1 {
			logger.done <- struct{}{}
		}

		logger.wg.Wait()
		logger.logfiledescriptor.Write([]byte(formatMessage(msgPrefixInfo, "STOPPED")))
		logger.close()
	}

	logger.mu.Unlock()
}

func (logger *FileLog) RenameOldLogFiles() {

	newestLogFilename := filepath.Join(logger.location, logger.filename)
	logFilename1 := filepath.Join(logger.location, logger.filename+".1")
	logFilename2 := filepath.Join(logger.location, logger.filename+".2")
	logFilename3 := filepath.Join(logger.location, logger.filename+".3")
	logFilename4 := filepath.Join(logger.location, logger.filename+".4")
	logFilename5 := filepath.Join(logger.location, logger.filename+".5")
	logFilename6 := filepath.Join(logger.location, logger.filename+".6")
	logFilename7 := filepath.Join(logger.location, logger.filename+".7")
	logFilename8 := filepath.Join(logger.location, logger.filename+".8")
	logFilename9 := filepath.Join(logger.location, logger.filename+".9")
	oldestLogFilename := filepath.Join(logger.location, logger.filename+".10")

	if _, err := os.Stat(logFilename9); err == nil {
		os.Rename(logFilename9, oldestLogFilename)
	}
	if _, err := os.Stat(logFilename8); err == nil {
		os.Rename(logFilename8, logFilename9)
	}
	if _, err := os.Stat(logFilename7); err == nil {
		os.Rename(logFilename7, logFilename8)
	}
	if _, err := os.Stat(logFilename6); err == nil {
		os.Rename(logFilename6, logFilename7)
	}
	if _, err := os.Stat(logFilename5); err == nil {
		os.Rename(logFilename5, logFilename6)
	}
	if _, err := os.Stat(logFilename4); err == nil {
		os.Rename(logFilename4, logFilename5)
	}
	if _, err := os.Stat(logFilename3); err == nil {
		os.Rename(logFilename3, logFilename4)
	}
	if _, err := os.Stat(logFilename2); err == nil {
		os.Rename(logFilename2, logFilename3)
	}
	if _, err := os.Stat(logFilename1); err == nil {
		os.Rename(logFilename1, logFilename2)
	}
	if _, err := os.Stat(newestLogFilename); err == nil {
		os.Rename(newestLogFilename, logFilename1)
	}
}

func (logger *FileLog) Rotate() {

	logger.mu.Lock()

	logfilename := filepath.Join(logger.location, logger.filename)

	if file, err := os.Stat(logfilename); err == nil {
		size := file.Size()
		if size > int64(logger.logsize) {
			if len(logger.rotate) < 1 {
				logger.rotate <- struct{}{}
			}
		}
	}

	logger.mu.Unlock()
}

func (logger *FileLog) bufferMessage(message string) {
	if (logger.currentsize.Load() < int32(logger.buffersize)) && logger.receive {
		logger.currentsize.Add(1)
		logger.buffer <- message
	}
}

func (logger *FileLog) flusher() {
	for {
		select {
		case <-logger.done:
			for {
				select {
				case msg := <-logger.buffer:
					logger.currentsize.Add(-1)
					logger.logfiledescriptor.Write([]byte(msg))
				default:
					return
				}
			}
		case <-logger.rotate:
			logger.close()
			logger.RenameOldLogFiles()
			logger.open()
		case msg := <-logger.buffer:
			logger.currentsize.Add(-1)
			logger.logfiledescriptor.Write([]byte(msg))
		}
	}
}

func (logger *FileLog) LogMessage(level int, format string, args ...interface{}) {

	if logger.loglevel >= level {

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

		logMessage := formatMessage(prefix, fmt.Sprintf(format, args...))

		logger.bufferMessage(logMessage)
	}
}

func (logger *FileLog) QuickLog(level int, format string, args ...interface{}) {

	logger.mu.Lock()
	if logger.loglevel >= level {

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

		logMessage := formatMessage(prefix, fmt.Sprintf(format, args...))

		err := logger.open()

		if err == nil {
			logger.logfiledescriptor.Write([]byte(logMessage))
		}

		logger.close()
	}

	logger.mu.Unlock()
}

func (logger *FileLog) PanicQuickLog(format string, args ...interface{}) {
	logger.mu.Lock()
	logMessage := formatMessage(msgPrefixPanic, fmt.Sprintf(format, args...))
	err := os.MkdirAll(logger.location, 0777)
	if err == nil {
		os.WriteFile(filepath.Join(logger.location, "panic."+logger.filename), []byte(logMessage), 0666)
	}
	logger.mu.Unlock()
}

func (logger *FileLog) InfoLog(format string, args ...interface{}) {
	if logger.loglevel >= Info {
		logMessage := formatMessage(msgPrefixInfo, fmt.Sprintf(format, args...))
		logger.bufferMessage(logMessage)
	}
}

func (logger *FileLog) DebugLog(format string, args ...interface{}) {
	if logger.loglevel >= Debug {
		logMessage := formatMessage(msgPrefixDebug, fmt.Sprintf(format, args...))
		logger.bufferMessage(logMessage)
	}
}

func (logger *FileLog) WarningLog(format string, args ...interface{}) {
	if logger.loglevel >= Warning {
		logMessage := formatMessage(msgPrefixWarning, fmt.Sprintf(format, args...))
		logger.bufferMessage(logMessage)
	}
}

func (logger *FileLog) TraceLog(format string, args ...interface{}) {
	if logger.loglevel >= Trace {
		logMessage := formatMessage(msgPrefixTrace, fmt.Sprintf(format, args...))
		logger.bufferMessage(logMessage)
	}
}

package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type FileLog struct {
	FileLogConfig
	receive           bool
	done              chan struct{}
	rotateSignal      chan struct{}
	queue             chan string
	queueSize         atomic.Int32
	logfiledescriptor *os.File
	wg                sync.WaitGroup
	mu                sync.Mutex
	bytesWritten      int
}

type FileLogConfig struct {
	filename string
	location string
	level    Level
	size     int
	buffer   int
	maxFiles int
}

type FileLogOptions func(*FileLogConfig)

func WithFilename(filename string) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.filename = filename
	}
}

func WithLocation(location string) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.location = location
	}
}

func WithLevel(level Level) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.level = level
	}
}

func WithSize(size int) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.size = size
	}
}

func WithBuffer(buffer int) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.buffer = buffer
	}
}

func WithMaxFiles(maxFiles int) FileLogOptions {
	return func(cfg *FileLogConfig) {
		cfg.maxFiles = maxFiles
	}
}

func NewFileLogConfig(options ...FileLogOptions) FileLogConfig {
	config := FileLogConfig{
		level:    NoLog,
		size:     defaultSize,
		buffer:   defaultBuffer,
		maxFiles: 10,
	}

	for _, option := range options {
		option(&config)
	}

	return config
}

func NewFileLogger(cfg FileLogConfig) *FileLog {

	logger := &FileLog{
		FileLogConfig: cfg,
		queue:         make(chan string, cfg.buffer),
		done:          make(chan struct{}, 1),
		rotateSignal:  make(chan struct{}, 1),
	}

	return logger
}

func (logger *FileLog) Update(cfg FileLogConfig) error {

	if (logger.filename != cfg.filename) ||
		(logger.location != cfg.location) ||
		(logger.level != cfg.level) ||
		(logger.size != cfg.size) ||
		(logger.buffer != cfg.buffer) ||
		(logger.maxFiles != cfg.maxFiles) {

		logger.Stop()

		logger.FileLogConfig = cfg
		logger.queue = make(chan string, cfg.buffer)

		return logger.Start()
	}

	return nil
}

func (logger *FileLog) open() error {
	var err error

	err = os.MkdirAll(logger.location, 0755)
	if err != nil {
		return err
	}

	logfilename := filepath.Join(logger.location, logger.filename)

	logger.logfiledescriptor, err = os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return err
}

func (logger *FileLog) close() error {
	return logger.logfiledescriptor.Close()
}

func (logger *FileLog) Start() error {

	logger.mu.Lock()
	defer logger.mu.Unlock()

	if logger.receive {
		return fmt.Errorf("logger already in running state")
	}

	err := logger.open()

	if err == nil {
		logger.logfiledescriptor.Write([]byte(formatMessage(info, "STARTED")))
		logger.receive = true
		logger.wg.Add(1)
		go func() {
			logger.flusher()
			logger.wg.Done()
		}()
	}

	return err
}

func (logger *FileLog) Stop() {

	logger.mu.Lock()
	defer logger.mu.Unlock()

	if logger.receive {
		logger.receive = false

		if len(logger.done) < 1 {
			logger.done <- struct{}{}
		}

		logger.wg.Wait()
		logger.logfiledescriptor.Write([]byte(formatMessage(info, "STOPPED")))
		logger.close()
	}

}

func (logger *FileLog) renameOldLogFiles() {

	newestLogFilename := filepath.Join(logger.location, logger.filename)
	var oldLogFilename, newLogFilename string

	for i := logger.maxFiles - 1; i > 0; i-- {
		oldLogFilename = filepath.Join(logger.location, fmt.Sprintf("%s.%d", logger.filename, i))
		newLogFilename = filepath.Join(logger.location, fmt.Sprintf("%s.%d", logger.filename, i+1))
		if _, err := os.Stat(oldLogFilename); err == nil {
			os.Rename(oldLogFilename, newLogFilename)
		}
	}

	if _, err := os.Stat(newestLogFilename); err == nil {
		os.Rename(newestLogFilename, filepath.Join(logger.location, logger.filename+".1"))
	}
}

func (logger *FileLog) Rotate() {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	if len(logger.rotateSignal) < 1 { // accept signal only if previous rotate signal is consumed
		logger.rotateSignal <- struct{}{}
	}
}

func (logger *FileLog) queueMessage(message string) {
	if (logger.queueSize.Load() < int32(logger.buffer)) && logger.receive {
		logger.queueSize.Add(1)
		logger.queue <- message
	}
}

func (logger *FileLog) write(msg []byte) {
	n, _ := logger.logfiledescriptor.Write(msg)
	logger.bytesWritten += n
	if logger.bytesWritten > logger.size {
		logger.Rotate()
		logger.bytesWritten = 0
	}
}

func (logger *FileLog) flusher() {
	for {
		select {
		case <-logger.done:
			for {
				select {
				case msg := <-logger.queue:
					logger.queueSize.Add(-1)
					logger.write([]byte(msg))
				default:
					return
				}
			}
		case <-logger.rotateSignal:
			logger.close()
			logger.renameOldLogFiles()
			logger.open()
		case msg := <-logger.queue:
			logger.queueSize.Add(-1)
			logger.write([]byte(msg))
		}
	}
}

func (logger *FileLog) WithLevel(level Level, format string, args ...interface{}) {

	if logger.level >= level {

		logMessage := formatMessage(level.String(), fmt.Sprintf(format, args...))

		logger.queueMessage(logMessage)
	}
}

func (logger *FileLog) Quick(level Level, format string, args ...interface{}) {

	logger.mu.Lock()
	defer logger.mu.Unlock()

	if logger.level >= level {

		logMessage := formatMessage(level.String(), fmt.Sprintf(format, args...))

		err := logger.open()

		if err == nil {
			logger.logfiledescriptor.Write([]byte(logMessage))
		}

		logger.close()
	}

}

func (logger *FileLog) Panic(format string, args ...interface{}) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logMessage := formatMessage(panic, fmt.Sprintf(format, args...))
	err := os.MkdirAll(logger.location, 0755)
	if err == nil {
		os.WriteFile(filepath.Join(logger.location, "panic."+logger.filename), []byte(logMessage), 0755)
	}
}

func (logger *FileLog) Info(format string, args ...interface{}) {
	if logger.level >= Info {
		logMessage := formatMessage(info, fmt.Sprintf(format, args...))
		logger.queueMessage(logMessage)
	}
}

func (logger *FileLog) Debug(format string, args ...interface{}) {
	if logger.level >= Debug {
		logMessage := formatMessage(debug, fmt.Sprintf(format, args...))
		logger.queueMessage(logMessage)
	}
}

func (logger *FileLog) Warning(format string, args ...interface{}) {
	if logger.level >= Warning {
		logMessage := formatMessage(warning, fmt.Sprintf(format, args...))
		logger.queueMessage(logMessage)
	}
}

func (logger *FileLog) Trace(format string, args ...interface{}) {
	if logger.level >= Trace {
		logMessage := formatMessage(trace, fmt.Sprintf(format, args...))
		logger.queueMessage(logMessage)
	}
}

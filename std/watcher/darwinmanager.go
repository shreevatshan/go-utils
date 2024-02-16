//go:build darwin
// +build darwin

package watcher

import (
	"dataexporter/pkg/std/log"
	"os"
	"os/signal"
	"syscall"
)

type service interface {
	Shutdown()
	GetName() string
	GetLogger() log.Logger
}

var s service

func signalHandler(signalChannel <-chan os.Signal) {
loop:
	for {
		signal := <-signalChannel
		switch signal {
		case syscall.SIGTERM:
			s.GetLogger().LogMessage(log.Info, "Control request received [%s]", signal.String())
			s.Shutdown()
			break loop
		default:
			s.GetLogger().LogMessage(log.Warning, "Unexpected control request [%s]", signal.String())
		}
	}
}

func StartServiceManager(ser service) {

	s = ser

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	signalHandler(signalChannel)
}

//go:build linux
// +build linux

package watcher

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/shreevatshan/go-utils/std/log"
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
			s.GetLogger().Info("Control request received [%s]", signal.String())
			s.Shutdown()
			break loop
		default:
			s.GetLogger().Info("Unexpected control request [%s]", signal.String())
		}
	}
}

func StartServiceManager(ser service) {

	s = ser

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	signalHandler(signalChannel)
}

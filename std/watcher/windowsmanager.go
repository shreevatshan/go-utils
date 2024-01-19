//go:build windows
// +build windows

package watcher

import (
	"dataexporter/pkg/std/log"
	"fmt"
	"strings"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

type service interface {
	Shutdown()
	GetName() string
	GetLogger() log.Logger
}

var s service

type WindowssService struct{}

var elog debug.Log

func (m *WindowssService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				output := strings.Join(args, "-")
				output += fmt.Sprintf("-%d", c.Context)
				elog.Info(1, output)
				s.Shutdown()
				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func runWindowsService() {

	var err error
	elog, err = eventlog.Open(s.GetName())
	if err != nil {
		return
	}

	defer elog.Close()

	elog.Info(1, fmt.Sprintf("Starting %s service", s.GetName()))
	run := svc.Run
	err = run(s.GetName(), &WindowssService{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", s.GetName(), err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", s.GetName()))
}

func StartServiceManager(ser service) {

	s = ser

	isWindowsservice, err := svc.IsWindowsService()
	if err != nil {
		s.GetLogger().LogMessage(log.Warning, "Failed to determine if running as service [%v]", err)
	}
	if isWindowsservice {
		runWindowsService()
	}
}

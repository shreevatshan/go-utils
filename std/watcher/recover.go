package watcher

import (
	"dataexporter/pkg/std/log"
	"runtime/debug"
	"time"
)

func RecoverOnPanic(f func(), opt ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			// Wait before retrying function execution
			time.Sleep(1 * time.Second)

			if len(opt) > 0 {
				logger := opt[0].(log.Logger)
				logger.PanicQuickLog("retrying function execution, panic [%v] triggered at\n%v", r, string(debug.Stack()))
				RecoverOnPanic(f, logger)
			} else {
				RecoverOnPanic(f)
			}
		}
	}()
	f()
}

func ExecuteSafe(f func(), opt ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			if len(opt) > 0 {
				logger := opt[0].(log.Logger)
				logger.PanicQuickLog("function execution stopped, panic [%v] triggered at\n%v", r, string(debug.Stack()))
			}
		}
	}()
	f()
}

/*
always call HandlePanic as a deferred function above the function which is susceptible to panicking
*/
func HandlePanic(opt ...interface{}) {
	if r := recover(); r != nil {
		if len(opt) > 0 {
			logger := opt[0].(log.Logger)
			logger.PanicQuickLog("function execution stopped, panic [%v] triggered at\n%v", r, string(debug.Stack()))
		}
	}
}

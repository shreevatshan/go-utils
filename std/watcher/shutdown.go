package watcher

import "sync"

type ShutdownHandler struct {
	listenerlist map[string]*ShutdownListener
	mu           sync.Mutex
}

type ShutdownListener struct {
	State chan struct{}
}

func InitShutdownHandler() *ShutdownHandler {
	shutdownhandler := &ShutdownHandler{
		listenerlist: make(map[string]*ShutdownListener),
	}
	return shutdownhandler
}

func (shutdownhandler *ShutdownHandler) AddListener(listenerName string, shutdownlistener *ShutdownListener) {
	shutdownhandler.mu.Lock()
	shutdownhandler.listenerlist[listenerName] = shutdownlistener
	shutdownhandler.mu.Unlock()
}

func (shutdownhandler *ShutdownHandler) RemoveListener(listenerName string) {
	shutdownhandler.mu.Lock()
	delete(shutdownhandler.listenerlist, listenerName)
	shutdownhandler.mu.Unlock()
}

func (shutdownhandler *ShutdownHandler) NotifyListener(listenerName string) {

	if _, exists := shutdownhandler.listenerlist[listenerName]; exists {
		shutdownlistener := shutdownhandler.listenerlist[listenerName]
		shutdownlistener.Shutdown()
	}
}

func (shutdownhandler *ShutdownHandler) NotifyAllListeners() {

	for _, shutdownlistener := range shutdownhandler.listenerlist {
		shutdownlistener.Shutdown()
	}
}

func (shutdownhandler *ShutdownHandler) GetListener(listenerName string) *ShutdownListener {
	shutdownlistener := &ShutdownListener{
		State: make(chan struct{}),
	}
	shutdownhandler.AddListener(listenerName, shutdownlistener)
	return shutdownlistener
}

func (shutdownlistener *ShutdownListener) Shutdown() {
	close(shutdownlistener.State)
}

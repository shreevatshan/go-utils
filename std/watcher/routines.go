package watcher

import (
	"fmt"
	"sync/atomic"
)

type RoutineManager struct {
	maxWorkers  int32
	workers     atomic.Int32
	managerName string
}

type Job interface {
	DoJob()
	QuitJob()
}

func (routineManager *RoutineManager) incrementWorker() {
	routineManager.workers.Add(1)
}

func (routineManager *RoutineManager) decrementWorker() {
	routineManager.workers.Add(-1)
}

func (routineManager *RoutineManager) getCurrentWorkers() int32 {
	return routineManager.workers.Load()
}

func (routineManager *RoutineManager) Execute(j Job) error {
	if routineManager.getCurrentWorkers() < routineManager.maxWorkers {
		routineManager.incrementWorker()
		go func() {
			defer routineManager.decrementWorker()
			ExecuteSafe(j.DoJob)
		}()
	} else {
		ExecuteSafe(j.QuitJob)
		return fmt.Errorf("all workers under %s routine manager are busy", routineManager.managerName)
	}

	return nil
}

func InitRoutineManager(managerName string, maxWorkers int32) *RoutineManager {
	routineManager := &RoutineManager{
		managerName: managerName,
		maxWorkers:  maxWorkers,
	}
	return routineManager
}

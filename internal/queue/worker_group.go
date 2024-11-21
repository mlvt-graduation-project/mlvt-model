package queue

import (
	"sync"
)

// WorkerGroup manages a group of workers.
type WorkerGroup struct {
	wg sync.WaitGroup
}

// NewWorkerGroup initializes a new WorkerGroup.
func NewWorkerGroup() *WorkerGroup {
	return &WorkerGroup{
		wg: sync.WaitGroup{},
	}
}

// AddWorker adds a worker's WaitGroup to the group.
func (wg *WorkerGroup) AddWorker(worker *Worker) {
	wg.wg.Add(1)
}

// WorkerGroupWait waits for all workers in the group to finish.
func (wg *WorkerGroup) Wait() {
	wg.wg.Wait()
}
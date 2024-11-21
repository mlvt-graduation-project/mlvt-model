package queue

import (
	"mlvt-api/api/model"
)

// JobQueue represents a job queue.
type JobQueue struct {
	queue chan *model.Job
	wg    *WorkerGroup
}

// NewJobQueue initializes a new JobQueue with a specified buffer size.
func NewJobQueue(bufferSize int) *JobQueue {
	return &JobQueue{
		queue: make(chan *model.Job, bufferSize),
		wg:    NewWorkerGroup(),
	}
}

// Enqueue adds a job to the queue.
func (jq *JobQueue) Enqueue(job *model.Job) {
	jq.queue <- job
}

// StartWorkers starts the specified number of worker goroutines.
func (jq *JobQueue) StartWorkers(numWorkers int, store *model.JobStatusStore, callbackURL string) {
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, jq.queue, store, callbackURL)
		jq.wg.AddWorker(worker)
		worker.Start()
	}
}

// Wait waits for all workers to finish.
func (jq *JobQueue) Wait() {
	jq.wg.Wait()
}

// Close gracefully shuts down the job queue by closing the queue channel.
func (jq *JobQueue) Close() {
	close(jq.queue)
}

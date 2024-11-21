package queue

import (
	"log"
	"mlvt-api/api/model"
)

// Worker represents a single worker that processes jobs.
type Worker struct {
	id          int
	jobQueue    <-chan *model.Job
	store       *model.JobStatusStore
	callbackURL string
	processor   *Processor
}

// NewWorker initializes a new Worker.
func NewWorker(id int, jobQueue <-chan *model.Job, store *model.JobStatusStore, callbackURL string) *Worker {
	return &Worker{
		id:          id,
		jobQueue:    jobQueue,
		store:       store,
		callbackURL: callbackURL,
		processor:   NewProcessor(),
	}
}

// Start begins the worker's job processing loop.
func (w *Worker) Start() {
	go func() {
		log.Printf("Worker %d started\n", w.id)
		for job := range w.jobQueue {
			log.Printf("Worker %d processing job ID: %s\n", w.id, job.ID)
			w.processJob(job)
		}
		log.Printf("Worker %d stopped\n", w.id)
	}()
}

// processJob handles the processing of a single job.
func (w *Worker) processJob(job *model.Job) {
	// Update job status to "in_progress"
	err := w.store.UpdateJob(job.ID, "in_progress", "")
	if err != nil {
		log.Printf("Worker %d: Failed to update job status to in_progress for job ID %s: %v", w.id, job.ID, err)
		return
	}

	// Process the job based on its type
	procErr := w.processor.Process(job)

	// Update job status based on processing result
	if procErr != nil {
		log.Printf("Worker %d: Job ID %s failed: %v", w.id, job.ID, procErr)
		updateErr := w.store.UpdateJob(job.ID, "failed", procErr.Error())
		if updateErr != nil {
			log.Printf("Worker %d: Failed to update job status to failed for job ID %s: %v", w.id, job.ID, updateErr)
		}
	} else {
		log.Printf("Worker %d: Job ID %s succeeded", w.id, job.ID)
		updateErr := w.store.UpdateJob(job.ID, "succeeded", "")
		if updateErr != nil {
			log.Printf("Worker %d: Failed to update job status to succeeded for job ID %s: %v", w.id, job.ID, updateErr)
		}
	}

	// Call the callback API to update job status
	cbErr := w.processor.CallCallbackAPI(w.callbackURL, job)
	if cbErr != nil {
		log.Printf("Worker %d: Failed to call callback API for job ID %s: %v", w.id, job.ID, cbErr)
		// Optionally, implement retry logic or log for manual intervention
	}
}

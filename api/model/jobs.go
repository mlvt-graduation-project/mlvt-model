package model

import (
	"errors"
	"sync"
)

// Job represents a processing job.
type Job struct {
	ID        string
	Type      string // "stt", "tts", "ttt"
	Request   interface{}
	Status    string // "succeeded", "failed"
	Error     string
	CreatedAt string
	UpdatedAt string
	Result    interface{}   // Optional: to store job result
	Done      chan struct{} // Channel to signal job completion
}

// JobStatusStore manages job statuses.
type JobStatusStore struct {
	sync.RWMutex
	Jobs map[string]*Job
}

// NewJobStatusStore initializes a new JobStatusStore.
func NewJobStatusStore() *JobStatusStore {
	return &JobStatusStore{
		Jobs: make(map[string]*Job),
	}
}

// AddJob adds a new job to the store.
func (s *JobStatusStore) AddJob(job *Job) {
	s.Lock()
	defer s.Unlock()
	s.Jobs[job.ID] = job
}

// UpdateJob updates the status, error, and optionally the result of a job.
func (s *JobStatusStore) UpdateJob(id, status, errMsg string, result interface{}) error {
	s.Lock()
	defer s.Unlock()
	job, exists := s.Jobs[id]
	if !exists {
		return errors.New("job not found")
	}
	job.Status = status
	job.Error = errMsg
	if result != nil {
		job.Result = result
	}
	return nil
}

// GetJob retrieves a job by ID.
func (s *JobStatusStore) GetJob(id string) (*Job, error) {
	s.RLock()
	defer s.RUnlock()
	job, exists := s.Jobs[id]
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}

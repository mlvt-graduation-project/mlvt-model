package handler

import (
	"mlvt-api/api/model"
	"mlvt-api/internal/queue"
)

// Handler encapsulates the dependencies for handlers.
type Handler struct {
	JobStore    *model.JobStatusStore
	JobQueue    *queue.JobQueue
	CallbackURL string
}

// NewHandler initializes a new Handler with dependencies.
func NewHandler(store *model.JobStatusStore, queue *queue.JobQueue, callbackURL string) *Handler {
	return &Handler{
		JobStore:    store,
		JobQueue:    queue,
		CallbackURL: callbackURL,
	}
}

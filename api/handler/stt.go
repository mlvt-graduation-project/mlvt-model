package handler

import (
	"log"
	"net/http"

	"mlvt-api/api/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// STTHandler handles the Speech-to-Text (STT) processing requests asynchronously.
func (h *Handler) STTHandler(c *gin.Context) {
	var req model.STTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Incoming STT request: %+v\n", req)

	// Generate a unique job ID
	jobID := uuid.New().String()

	// Create a new job
	job := &model.Job{
		ID:      jobID,
		Type:    "stt",
		Request: &req,
		Status:  "received",
	}

	// Add job to the status store
	h.JobStore.AddJob(job)

	// Enqueue the job for background processing
	h.JobQueue.Enqueue(job)

	// Respond immediately with the job ID and status
	c.JSON(http.StatusAccepted, gin.H{
		"message": "STT processing request received",
		"job_id":  jobID,
		"status":  job.Status,
	})
}

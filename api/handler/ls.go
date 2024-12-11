package handler

import (
	"log"
	"mlvt-api/api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LSHandler handles the Lip-Sync (LS) processing requests synchronously.
func (h *Handler) LSHandler(c *gin.Context) {
	var req model.LSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Incoming LS request: %+v\n", req)

	// Generate a unique job ID
	jobID := uuid.New().String()

	// Create a new job
	job := &model.Job{
		ID:        jobID,
		Type:      "ls",
		Request:   &req,
		Status:    "received",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	// Add job to the status store
	h.JobStore.AddJob(job)

	// Enqueue the job for processing
	h.JobQueue.Enqueue(job)

	// Wait for the job to be processed with a timeout
	select {
	case <-job.Done:
		// Job completed
		if job.Status == "succeeded" {
			c.JSON(http.StatusOK, gin.H{
				"message": "LS processing completed",
				"job_id":  job.ID,
				"status":  job.Status,
				"result":  job.Result, // Include the result in the response if available
			})
		} else {
			// Job failed
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "LS processing failed",
				"job_id":  job.ID,
				"status":  job.Status,
				"error":   job.Error,
			})
		}
	case <-time.After(5 * time.Minute):
		// Timeout after 5 minutes
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"message": "LS processing timed out",
			"job_id":  job.ID,
			"status":  "timeout",
		})
		return
	}
}

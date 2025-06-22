package handler

import (
	"log"
	"net/http"
	"time"

	"mlvt-api/api/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TTTHandler handles the Text-to-Text (TTT) processing requests synchronously.
func (h *Handler) TTTHandler(c *gin.Context) {
	var req model.TTTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Incoming TTT request: %+v\n", req)

	// Generate a unique job ID
	jobID := uuid.New().String()
	startTime := time.Now()

	// Notify request received
	h.notifyRequest(c, "TTT", jobID)

	// Create a new job
	job := &model.Job{
		ID:        jobID,
		Type:      "ttt",
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
		processingTime := time.Since(startTime)
		// Job completed successfully
		if job.Status == "succeeded" {
			h.notifySuccess("TTT", job.ID, processingTime)
			c.JSON(http.StatusOK, gin.H{
				"message": "TTT processing completed",
				"job_id":  job.ID,
				"status":  job.Status,
				"result":  job.Result, // Include the result in the response if available
			})
		} else {
			// Job failed
			errorMsg := "Unknown error"
			if job.Error != "" {
				errorMsg = job.Error
			}
			h.notifyFailure("TTT", job.ID, errorMsg, processingTime)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "TTT processing failed",
				"job_id":  job.ID,
				"status":  job.Status,
				"error":   job.Error,
			})
		}
	case <-time.After(5 * time.Minute):
		// Timeout after 5 minutes
		h.notifyTimeout("TTT", job.ID)
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"message": "TTT processing timed out",
			"job_id":  job.ID,
			"status":  "timeout",
		})
		return
	}
}

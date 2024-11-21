// handler/tts_handler.go
package handler

import (
	"log"
	"net/http"

	"mlvt-api/api/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) TTSHandler(c *gin.Context) {
	var req model.TTSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Incoming TTS request: %+v\n", req)

	// Generate a unique job ID
	jobID := uuid.New().String()

	// Create a new job
	job := &model.Job{
		ID:      jobID,
		Type:    "tts",
		Request: &req,
		Status:  "received",
	}

	// Add job to the status store
	h.JobStore.AddJob(job)

	// Enqueue the job for background processing
	h.JobQueue.Enqueue(job)

	// Respond immediately with the job ID and status
	c.JSON(http.StatusAccepted, gin.H{
		"message": "TTS processing request received",
		"job_id":  jobID,
		"status":  job.Status,
	})
}

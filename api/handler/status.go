package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusHandler handles requests to check the status of a job.
func (h *Handler) StatusHandler(c *gin.Context) {
	jobID := c.Param("job_id")
	job, err := h.JobStore.GetJob(jobID)
	if err != nil {
		log.Printf("Status check failed for job ID %s: %v", jobID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job_id": job.ID,
		"type":   job.Type,
		"status": job.Status,
		"error":  job.Error,
	})
}

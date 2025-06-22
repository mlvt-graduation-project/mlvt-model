package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatusHandler handles requests to check the status of a job.
func (h *Handler) StatusHandler(c *gin.Context) {
	jobID := c.Param("job_id")

	// Notify status check request
	clientIP := c.ClientIP()
	message := fmt.Sprintf("📊 Status Check\n"+
		"Job ID: %s\n"+
		"Client IP: %s\n"+
		"Time: %s",
		jobID, clientIP, time.Now().Format("15:04:05"))
	h.notifyTelegram(message)

	job, err := h.JobStore.GetJob(jobID)
	if err != nil {
		log.Printf("Status check failed for job ID %s: %v", jobID, err)
		// Notify job not found
		notFoundMsg := fmt.Sprintf("❌ Status Check Failed\n"+
			"Job ID: %s\n"+
			"Error: Job not found\n"+
			"Client IP: %s\n"+
			"Time: %s",
			jobID, clientIP, time.Now().Format("15:04:05"))
		h.notifyTelegram(notFoundMsg)
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	response := gin.H{
		"job_id": job.ID,
		"type":   job.Type,
		"status": job.Status,
		"error":  job.Error,
	}

	// Optionally include the result if available
	if job.Result != nil {
		response["result"] = job.Result
	}

	// Notify successful status check
	statusMsg := fmt.Sprintf("✅ Status Check Success\n"+
		"Job ID: %s\n"+
		"Job Type: %s\n"+
		"Status: %s\n"+
		"Client IP: %s\n"+
		"Time: %s",
		job.ID, job.Type, job.Status, clientIP, time.Now().Format("15:04:05"))
	h.notifyTelegram(statusMsg)

	c.JSON(http.StatusOK, response)
}

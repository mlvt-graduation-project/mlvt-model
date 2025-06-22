package handler

import (
	"fmt"
	"log"
	"mlvt-api/api/model"
	"mlvt-api/internal/queue"
	"mlvt-api/notify"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler encapsulates the dependencies for handlers.
type Handler struct {
	JobStore *model.JobStatusStore
	JobQueue *queue.JobQueue
}

// NewHandler initializes a new Handler with dependencies.
func NewHandler(store *model.JobStatusStore, queue *queue.JobQueue) *Handler {
	return &Handler{
		JobStore: store,
		JobQueue: queue,
	}
}

// notifyTelegram sends a notification to Telegram and logs any errors
func (h *Handler) notifyTelegram(message string) {
	if err := notify.SendTelegram(message); err != nil {
		log.Printf("Failed to send Telegram notification: %v", err)
	}
}

// notifyRequest sends a notification when a request is received
func (h *Handler) notifyRequest(c *gin.Context, jobType string, jobID string) {
	clientIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	message := fmt.Sprintf("🔔 New %s Request\n"+
		"Job ID: %s\n"+
		"Client IP: %s\n"+
		"User Agent: %s\n"+
		"Time: %s",
		jobType, jobID, clientIP, userAgent, time.Now().Format("15:04:05"))
	h.notifyTelegram(message)
}

// notifySuccess sends a notification when a request succeeds
func (h *Handler) notifySuccess(jobType string, jobID string, processingTime time.Duration) {
	message := fmt.Sprintf("✅ %s Success\n"+
		"Job ID: %s\n"+
		"Processing Time: %v\n"+
		"Completed: %s",
		jobType, jobID, processingTime, time.Now().Format("15:04:05"))
	h.notifyTelegram(message)
}

// notifyFailure sends a notification when a request fails
func (h *Handler) notifyFailure(jobType string, jobID string, errorMsg string, processingTime time.Duration) {
	message := fmt.Sprintf("❌ %s Failed\n"+
		"Job ID: %s\n"+
		"Error: %s\n"+
		"Processing Time: %v\n"+
		"Failed: %s",
		jobType, jobID, errorMsg, processingTime, time.Now().Format("15:04:05"))
	h.notifyTelegram(message)
}

// notifyTimeout sends a notification when a request times out
func (h *Handler) notifyTimeout(jobType string, jobID string) {
	message := fmt.Sprintf("⏰ %s Timeout\n"+
		"Job ID: %s\n"+
		"Timeout after 5 minutes\n"+
		"Time: %s",
		jobType, jobID, time.Now().Format("15:04:05"))
	h.notifyTelegram(message)
}

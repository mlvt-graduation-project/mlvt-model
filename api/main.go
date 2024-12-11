package main

import (
	"log"
	"mlvt-api/api/handler"
	"mlvt-api/api/model"
	"mlvt-api/internal/queue"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize JobStatusStore
	jobStore := model.NewJobStatusStore()

	// Initialize JobQueue with a buffer size of 100
	jobQueue := queue.NewJobQueue(100)

	// Start background workers (e.g., 5 workers)
	// Updated StartWorkers signature if CallbackURL is removed
	jobQueue.StartWorkers(5, jobStore, "") // Passing empty string or modify StartWorkers to remove the parameter

	// Initialize Handler with dependencies (no CallbackURL)
	h := handler.NewHandler(jobStore, jobQueue)

	router := gin.Default()

	// Register routes.
	router.POST("/stt", h.STTHandler)
	router.POST("/tts", h.TTSHandler)
	router.POST("/ttt", h.TTTHandler)
	router.POST("/ls", h.LSHandler)
	router.GET("/status/:job_id", h.StatusHandler)

	// Start the server in a goroutine
	go func() {
		if err := router.Run("0.0.0.0:8000"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	log.Println("Server started on port 8000")

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Close the job queue channel to stop workers gracefully
	jobQueue.Close()

	// Wait for all workers to finish
	jobQueue.Wait()

	log.Println("Server gracefully stopped")
}

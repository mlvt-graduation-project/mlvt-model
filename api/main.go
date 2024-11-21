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

	// Set the callback URL (replace with your actual callback API URL)
	// callbackURL := os.Getenv("CALLBACK_URL")
	// if callbackURL == "" {
	// 	log.Fatal("CALLBACK_URL environment variable not set")
	// }
	callbackURL := "temp"

	// Start background workers (e.g., 5 workers)
	jobQueue.StartWorkers(5, jobStore, callbackURL)

	// Initialize Handler with dependencies
	h := handler.NewHandler(jobStore, jobQueue, callbackURL)

	router := gin.Default()

	// Register routes.
	router.POST("/stt", h.STTHandler)
	router.POST("/tts", h.TTSHandler)
	router.POST("/ttt", h.TTTHandler)
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

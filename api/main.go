package main

import (
	"mlvt-api/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register routes.
	router.POST("/stt", handler.STTHandler)
	router.POST("/tts", handler.TTSHandler)
	router.POST("/ttt", handler.TTTHandler)

	// Start the server.
	router.Run("0.0.0.0:8000") // Adjust the port as needed.
}

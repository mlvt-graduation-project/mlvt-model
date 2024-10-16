package main

import (
	"mlvt-api/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/process-stt", handler.ProcessSTT)
	// router.POST("/process-ttt", handler.ProcessTTT)
	// router.POST("/process-tts", handler.ProcessTTS)
	// router.POST("/process-way2lip", handler.ProcessWay2Lip)

	// run server on port 8000
	router.Run(":8000")
}

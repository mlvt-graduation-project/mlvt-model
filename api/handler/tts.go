package handler

import (
	"mlvt-api/api/model"
	"mlvt-api/internal/command"
	"mlvt-api/internal/python"
	utils "mlvt-api/pkg"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func TTSHandler(c *gin.Context) {
	var req model.TTSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputFileName := filepath.Base(filepath.Clean(req.InputFileName))
	outputFileName := filepath.Base(filepath.Clean(req.OutputFileName))

	inputFilePath := filepath.Join("data", "input", "tts", inputFileName)
	outputFilePath := filepath.Join("data", "output", "tts", outputFileName)

	if err := utils.DownloadFile(req.InputLink, inputFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download input file"})
		return
	}

	if err := command.RunTTS(python.Py3_11, inputFilePath, outputFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute TTS script", "details": err.Error()})
		return
	}

	if err := utils.UploadFile(outputFilePath, req.OutputLink); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload output file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TTS processing completed successfully"})
}

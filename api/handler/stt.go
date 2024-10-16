package handler

import (
	"mlvt-api/api/model"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ProcessSTT(c *gin.Context) {
	var req model.STTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputFile := filepath.Join("..", "data", "input", "stt", req.InputFile)
	outputFile := filepath.Join("..", "data", "output", "stt", req.OutputFile)

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input file does not exist."})
		return
	}

	outputDir := filepath.Dir(outputFile)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create output directory."})
			return
		}
	}

	cmd := exec.Command("python3.11", "../scripts/stt.py", "--input_file", inputFile, "--output_file", outputFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": string(output)})
}

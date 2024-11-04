package handler

import (
	"log"
	"mlvt-api/api/model"
	"mlvt-api/internal/command"
	"mlvt-api/internal/python"
	utils "mlvt-api/pkg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func STTHandler(c *gin.Context) {
	var req model.STTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Incoming request: %+v\n", req)

	// Sanitize the original filenames.
	originalInputName := filepath.Base(filepath.Clean(req.InputFileName))
	originalOutputName := filepath.Base(filepath.Clean(req.OutputFileName))

	// Extract the file extension from the input filename
	inputExt := filepath.Ext(originalInputName)
	if inputExt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input file must have an extension"})
		return
	}

	// Define directories
	inputDir := filepath.Join("data", "input", "stt")
	outputDir := filepath.Join("data", "output", "stt")

	// Ensure directories exist
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		log.Printf("Failed to create input directory: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create input directory"})
		return
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Printf("Failed to create output directory: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create output directory"})
		return
	}

	// Define full paths
	originalInputPath := filepath.Join(inputDir, originalInputName)
	outputFilePath := filepath.Join(outputDir, originalOutputName)

	// Step 1: Download the input file with the original filename
	log.Printf("Downloading input file from %s to %s\n", req.InputLink, originalInputPath)
	if err := utils.DownloadFile(req.InputLink, originalInputPath); err != nil {
		log.Printf("Failed to download input file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download input file"})
		return
	}
	log.Printf("Successfully downloaded input file to %s\n", originalInputPath)

	// Step 3: Execute the STT script with the original filenames
	log.Printf("Executing STT script for %s\n", originalInputName)
	if err := command.RunSTT(python.Py3, originalInputName, originalOutputName); err != nil {
		log.Printf("Failed to execute STT script: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute STT script", "details": err.Error()})
		return
	}
	log.Printf("Successfully executed STT script for %s\n", originalInputName)

	// Step 4: Upload the output file with the original output filename
	originalOutputPath := filepath.Join(outputDir, originalOutputName)
	log.Printf("Renaming output file from %s to %s\n", outputFilePath, originalOutputPath)
	if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
		log.Printf("Failed to rename output file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rename output file"})
		return
	}
	log.Printf("Successfully renamed output file to %s\n", originalOutputPath)

	log.Printf("Uploading output file from %s to %s\n", originalOutputPath, req.OutputLink)
	if err := utils.UploadFile(originalOutputPath, req.OutputLink); err != nil {
		log.Printf("Failed to upload output file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload output file"})
		return
	}
	log.Printf("Successfully uploaded output file to %s\n", req.OutputLink)

	// Optional: Clean up the random-named input and output files after processing
	// if err := os.Remove(randomInputPath); err != nil {
	//     // Log the error, but don't fail the request
	//     // You can replace this with a proper logging mechanism
	//     println("Failed to remove input file:", err.Error())
	// }
	// if err := os.Remove(originalOutputPath); err != nil {
	//     println("Failed to remove output file:", err.Error())
	// }

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "STT processing completed successfully"})
}

// import (
// 	"crypto/rand"
// 	"errors"
// 	"mlvt-api/api/model"
// 	"mlvt-api/internal/command"
// 	"mlvt-api/internal/python"
// //	utils "mlvt-api/pkg"
// 	"net/http"
// 	"os"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// // characterSet defines the allowed characters for the random filename.
// const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// // generateRandomFilename generates a random filename with the specified length and extension.
// // It uses the defined character set to build the filename.
// func generateRandomFilename(length int, extension string) (string, error) {
// 	if length <= 0 {
// 		return "", errors.New("filename length must be greater than zero")
// 	}

// 	bytes := make([]byte, length)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}

// 	for i, b := range bytes {
// 		bytes[i] = characterSet[b%byte(len(characterSet))]
// 	}

// 	return string(bytes) + extension, nil
// }

// func STTHandler(c *gin.Context) {
// 	var req model.STTRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Sanitize the original filenames.
// 	originalInputName := filepath.Base(filepath.Clean(req.InputFileName))
// //	originalOutputName := filepath.Base(filepath.Clean(req.OutputFileName))

// 	// Extract the file extension from the input filename
// 	inputExt := filepath.Ext(originalInputName)
// 	if inputExt == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Input file must have an extension"})
// 		return
// 	}

// 	// Generate random filenames for input and output
// //	randomInputName, err := generateRandomFilename(5, inputExt)
// //	if err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate random input filename"})
// //		return
// //	}

// 	randomOutputName, err := generateRandomFilename(5, ".txt") // Assuming STT output is text
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate random output filename"})
// 		return
// 	}

// 	// Define directories
// 	inputDir := filepath.Join("data", "input", "stt")
// 	outputDir := filepath.Join("data", "output", "stt")

// 	// Ensure directories exist
// 	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create input directory"})
// 		return
// 	}
// 	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create output directory"})
// 		return
// 	}

// 	// Define full paths
// //	originalInputPath := filepath.Join(inputDir, originalInputName)
// //	randomInputPath := filepath.Join(inputDir, randomInputName)
// //	outputFilePath := filepath.Join(outputDir, randomOutputName)

// 	// Step 1: Download the input file with the original filename
// //	if err := utils.DownloadFile(req.InputLink, originalInputPath); err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download input file"})
// //		return
// //	}

// 	// Step 2: Rename the downloaded file to the random filename
// //	if err := os.Rename(originalInputPath, randomInputPath); err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rename input file"})
// //		return
// //	}

// 	// Step 3: Execute the STT script with the random filenames
// 	if err := command.RunSTT(python.Py3, originalInputName, randomOutputName); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute STT script", "details": err.Error()})
// 		return
// 	}

// 	// Step 4: Upload the output file with the original output filename
// 	// Rename the output file back to the original output name before uploading
// //	originalOutputPath := filepath.Join(outputDir, originalOutputName)
// //	if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rename output file"})
// //		return
// //	}

// //	if err := utils.UploadFile(originalOutputPath, req.OutputLink); err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload output file"})
// //		return
// //	}

// 	// Optional: Clean up the random-named input and output files after processing
// 	// if err := os.Remove(randomInputPath); err != nil {
// 	// 	// Log the error, but don't fail the request
// 	// 	// You can replace this with a proper logging mechanism
// 	// 	println("Failed to remove input file:", err.Error())
// 	// }
// 	// if err := os.Remove(originalOutputPath); err != nil {
// 	// 	println("Failed to remove output file:", err.Error())
// 	// }

// 	// Respond with success
// 	c.JSON(http.StatusOK, gin.H{"message": "STT processing completed successfully"})
// }

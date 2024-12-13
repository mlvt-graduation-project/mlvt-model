package queue

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"mlvt-api/api/model"
	"mlvt-api/internal/command"
	"mlvt-api/internal/python"
	utils "mlvt-api/pkg"
)

// Processor handles the processing of jobs and callback interactions.
type Processor struct{}

// NewProcessor initializes a new Processor.
func NewProcessor() *Processor {
	return &Processor{}
}

// Process handles the processing of a job based on its type.
func (p *Processor) Process(job *model.Job) error {
	switch job.Type {
	case "ttt":
		req, ok := job.Request.(*model.TTTRequest)
		if !ok {
			return model.ErrInvalidRequestType
		}
		return p.processTTT(job, req)
	case "tts":
		req, ok := job.Request.(*model.TTSRequest)
		if !ok {
			return model.ErrInvalidRequestType
		}
		return p.processTTS(job, req)
	case "stt":
		req, ok := job.Request.(*model.STTRequest)
		if !ok {
			return model.ErrInvalidRequestType
		}
		return p.processSTT(job, req)
	case "ls":
		req, ok := job.Request.(*model.LSRequest)
		if !ok {
			return model.ErrInvalidRequestType
		}
		return p.processLS(job, req)
	default:
		return model.ErrUnknownJobType
	}
}

// processTTT processes a Text-to-Text (TTT) job.
func (p *Processor) processTTT(job *model.Job, req *model.TTTRequest) error {
	// Define directories
	inputDir := filepath.Join("data", "input", "ttt")
	outputDir := filepath.Join("data", "output", "ttt")

	// Ensure directories exist
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create input directory: %v", err)
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Define full paths
	originalInputPath := filepath.Join(inputDir, req.InputFileName)
	outputFilePath := filepath.Join(outputDir, req.OutputFileName)

	// Step 1: Download the input file
	log.Printf("Job ID %s: Downloading input file from %s to %s", job.ID, req.InputLink, originalInputPath)
	if err := utils.DownloadFile(req.InputLink, originalInputPath); err != nil {
		return fmt.Errorf("failed to download input file: %v", err)
	}

	// Step 2: Execute the TTT script
	log.Printf("Job ID %s: Executing TTT script for %s using model %s", job.ID, req.InputFileName, req.Model)
	modelPath, err := req.GetModelPath()
	if err != nil {
		return fmt.Errorf("failed to get model path: %v", err)
	}

	if err := command.RunTTT(python.Py3, modelPath, req.InputFileName, req.OutputFileName, req.SourceLang, req.TargetLang); err != nil {
		return fmt.Errorf("failed to execute TTT script: %v", err)
	}

	// Step 3: Rename the output file if necessary
	originalOutputPath := filepath.Join(outputDir, req.OutputFileName)
	log.Printf("Job ID %s: Renaming output file from %s to %s", job.ID, outputFilePath, originalOutputPath)
	if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
		return fmt.Errorf("failed to rename output file: %v", err)
	}

	// Step 4: Upload the output file
	log.Printf("Job ID %s: Uploading output file from %s to %s", job.ID, originalOutputPath, req.OutputLink)
	if err := utils.UploadFile(originalOutputPath, req.OutputLink); err != nil {
		return fmt.Errorf("failed to upload output file: %v", err)
	}

	// Optionally, set the Result field
	job.Result = "TTT processing completed successfully."

	return nil
}

// processTTS processes a Text-to-Speech (TTS) job.
func (p *Processor) processTTS(job *model.Job, req *model.TTSRequest) error {
	// Define directories
	inputDir := filepath.Join("data", "input", "tts")
	outputDir := filepath.Join("data", "output", "tts")

	// Ensure directories exist
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create input directory: %v", err)
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Define full paths
	originalInputPath := filepath.Join(inputDir, req.InputFileName)
	outputFilePath := filepath.Join(outputDir, req.OutputFileName)

	// Step 1: Download the input file
	log.Printf("Job ID %s: Downloading input file from %s to %s", job.ID, req.InputLink, originalInputPath)
	if err := utils.DownloadFile(req.InputLink, originalInputPath); err != nil {
		return fmt.Errorf("failed to download input file: %v", err)
	}

	// Step 2: Execute the TTS script
	log.Printf("Job ID %s: Executing TTS script for %s using model %s", job.ID, req.InputFileName, req.Model)
	modelPath, err := req.GetModelPath()
	if err != nil {
		return fmt.Errorf("failed to get model path: %v", err)
	}

	if err := command.RunTTS(python.Py3, modelPath, req.InputFileName, req.OutputFileName); err != nil {
		return fmt.Errorf("failed to execute TTS script: %v", err)
	}

	// Step 3: Rename the output file if necessary
	originalOutputPath := filepath.Join(outputDir, req.OutputFileName)
	log.Printf("Job ID %s: Renaming output file from %s to %s", job.ID, outputFilePath, originalOutputPath)
	if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
		return fmt.Errorf("failed to rename output file: %v", err)
	}

	// Step 4: Upload the output file
	log.Printf("Job ID %s: Uploading output file from %s to %s", job.ID, originalOutputPath, req.OutputLink)
	if err := utils.UploadFile(originalOutputPath, req.OutputLink); err != nil {
		return fmt.Errorf("failed to upload output file: %v", err)
	}

	// Optionally, set the Result field
	job.Result = "TTS processing completed successfully."

	return nil
}

// processSTT processes a Speech-to-Text (STT) job.
func (p *Processor) processSTT(job *model.Job, req *model.STTRequest) error {
	// Define directories
	inputDir := filepath.Join("data", "input", "stt")
	outputDir := filepath.Join("data", "output", "stt")

	// Ensure directories exist
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create input directory: %v", err)
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Define full paths
	originalInputPath := filepath.Join(inputDir, req.InputFileName)
	outputFilePath := filepath.Join(outputDir, req.OutputFileName)

	// Step 1: Download the input file
	log.Printf("Job ID %s: Downloading input file from %s to %s", job.ID, req.InputLink, originalInputPath)
	if err := utils.DownloadFile(req.InputLink, originalInputPath); err != nil {
		return fmt.Errorf("failed to download input file: %v", err)
	}

	// Step 2: Execute the STT script
	log.Printf("Job ID %s: Executing STT script for %s using model %s", job.ID, req.InputFileName, req.Model)
	modelPath, err := req.GetModelPath()
	if err != nil {
		return fmt.Errorf("failed to get model path: %v", err)
	}

	if err := command.RunSTT(python.Py3, modelPath, req.InputFileName, req.OutputFileName); err != nil {
		return fmt.Errorf("failed to execute STT script: %v", err)
	}

	// Step 3: Rename the output file if necessary
	originalOutputPath := filepath.Join(outputDir, req.OutputFileName)
	log.Printf("Job ID %s: Renaming output file from %s to %s", job.ID, outputFilePath, originalOutputPath)
	if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
		return fmt.Errorf("failed to rename output file: %v", err)
	}

	// Step 4: Upload the output file
	log.Printf("Job ID %s: Uploading output file from %s to %s", job.ID, originalOutputPath, req.OutputLink)
	if err := utils.UploadFile(originalOutputPath, req.OutputLink); err != nil {
		return fmt.Errorf("failed to upload output file: %v", err)
	}

	// Optionally, set the Result field
	job.Result = "STT processing completed successfully."

	return nil
}

func (p *Processor) processLS(job *model.Job, req *model.LSRequest) error {
	// Define directories
	inputDir := filepath.Join("data", "input", "ls")
	outputDir := filepath.Join("data", "output", "ls")

	// Ensure directories exist
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create input directory: %v", err)
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Define full paths
	originalVideoInputPath := filepath.Join(inputDir, req.InputVideoFileName)
	originalAudioInputPath := filepath.Join(inputDir, req.InputAudioFileName)
	outputFilePath := filepath.Join(outputDir, req.OutputVideoFileName)

	// Step 1: Download the input files
	log.Printf("Job ID %s: Downloading video file from %s to %s", job.ID, req.InputVideoLink, originalVideoInputPath)
	if err := utils.DownloadFile(req.InputVideoLink, originalVideoInputPath); err != nil {
		return fmt.Errorf("failed to download video input file: %v", err)
	}

	log.Printf("Job ID %s: Downloading audio file from %s to %s", job.ID, req.InputAudioLink, originalAudioInputPath)
	if err := utils.DownloadFile(req.InputAudioLink, originalAudioInputPath); err != nil {
		return fmt.Errorf("failed to download audio input file: %v", err)
	}

	// Step 2: Execute the LS script
	// For now, we assume a command like: command.RunLS(python.Py3, modelPath, inputVideo, inputAudio, outputFile)
	// The modelPath may not be used if LS doesn't require it, you can pass an empty string or handle accordingly.
	log.Printf("Job ID %s: Executing LS script for %s and %s", job.ID, req.InputVideoFileName, req.InputAudioFileName)
	// If the LS script doesn't use model, you can omit or handle model related parameters.
	if err := command.RunLS(python.Py3, "", req.InputVideoFileName, req.InputAudioFileName, req.OutputVideoFileName); err != nil {
		return fmt.Errorf("failed to execute LS script: %v", err)
	}

	// Step 3: Rename the output file if necessary (depending on how the script outputs)
	originalOutputPath := filepath.Join(outputDir, req.OutputVideoFileName)
	if originalOutputPath != outputFilePath {
		log.Printf("Job ID %s: Renaming output file from %s to %s", job.ID, outputFilePath, originalOutputPath)
		if err := os.Rename(outputFilePath, originalOutputPath); err != nil {
			return fmt.Errorf("failed to rename output file: %v", err)
		}
	}

	// Step 4: Upload the output file
	log.Printf("Job ID %s: Uploading output file from %s to %s", job.ID, originalOutputPath, req.OutputVideoLink)
	if err := utils.UploadFile(originalOutputPath, req.OutputVideoLink); err != nil {
		return fmt.Errorf("failed to upload output file: %v", err)
	}

	// Optionally, set the Result field
	job.Result = "LS processing completed successfully."

	return nil
}

package model

import (
	"fmt"
	"path/filepath"
)

const (
	BaseScriptsDir = "/home/ec2-user/mlvt-api/scripts"
)

// Maps to hold model paths for each processing type.
var (
	STTModelPaths = map[ModelEntity]string{
		ModelSTT_Whisper: filepath.Join(BaseScriptsDir, "stt", "stt.py"),
	}

	TTTModelPaths = map[ModelEntity]string{
		ModelTTT_Gemini: filepath.Join(BaseScriptsDir, "ttt", "ttt.py"),
	}

	TTSModelPaths = map[ModelEntity]string{
		ModelTTS_LightSpeed: filepath.Join(BaseScriptsDir, "tts", "tts.py"),
		ModelTTS_YourTTS:    filepath.Join(BaseScriptsDir, "tts", "your_tts.py"),
		ModelTTS_XTTS:       filepath.Join(BaseScriptsDir, "tts", "xtts.py"),
	}
)

type ModelPathGetter interface {
	GetModelPath() (string, error)
}

// GetModelPath for STTRequest
func (r *STTRequest) GetModelPath() (string, error) {
	path, exists := STTModelPaths[r.Model]
	if !exists {
		// Return default model path if model not found
		path, exists = STTModelPaths[ModelSTT_Whisper]
		if !exists {
			return "", fmt.Errorf("default STT model path not defined")
		}
	}
	return path, nil
}

// GetModelPath for TTSRequest
func (r *TTSRequest) GetModelPath() (string, error) {
	path, exists := TTSModelPaths[r.Model]
	if !exists {
		// Return default model path if model not found
		path, exists = TTSModelPaths[ModelTTS_LightSpeed]
		if !exists {
			return "", fmt.Errorf("default TTS model path not defined")
		}
	}
	return path, nil
}

// GetModelPath for TTTRequest
func (r *TTTRequest) GetModelPath() (string, error) {
	path, exists := TTTModelPaths[r.Model]
	if !exists {
		// Return default model path if model not found
		path, exists = TTTModelPaths[ModelTTT_Gemini]
		if !exists {
			return "", fmt.Errorf("default TTT model path not defined")
		}
	}
	return path, nil
}

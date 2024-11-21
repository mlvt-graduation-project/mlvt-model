package model

type ModelEntity string

const (
	// Speech-to-Text models
	ModelSTT_Whisper ModelEntity = "whisper"

	// Text-to-Text models
	ModelTTT_Gemini ModelEntity = "gemini"

	// Text-to-Speech models
	ModelTTS_LightSpeed ModelEntity = "light-speed"
	ModelTTS_YourTTS    ModelEntity = "your-tts"
	ModelTTS_XTTS       ModelEntity = "xtts"

	// Voice Cloning models
	ModelVC_FreeVC ModelEntity = "free-vc"

	// Lipsync models
	ModelLS_Way2Lips ModelEntity = "way2lips"
)

package model

type BaseRequest struct {
	InputFileName  string      `json:"input_file_name" binding:"required"`
	InputLink      string      `json:"input_link" binding:"required"`
	OutputFileName string      `json:"output_file_name" binding:"required"`
	OutputLink     string      `json:"output_link" binding:"required"`
	Model          ModelEntity `json:"model"`
}

type BaseLang struct {
	SourceLang string `json:"source_language" binding:"required"`
	TargetLang string `json:"target_language" binding:"required"`
}

type STTRequest struct {
	BaseRequest
}

type TTSRequest struct {
	BaseRequest
	Lang string `json:"lang" binding:"required"`
}

type TTTRequest struct {
	BaseRequest
	BaseLang
}

type LSRequest struct {
	InputVideoFileName  string      `json:"input_video_file_name" binding:"required"`
	InputVideoLink      string      `json:"input_video_link" binding:"required"`
	InputAudioFileName  string      `json:"input_audio_file_name" binding:"required"`
	InputAudioLink      string      `json:"input_audio_link" binding:"required"`
	OutputVideoFileName string      `json:"output_video_file_name" binding:"required"`
	OutputVideoLink     string      `json:"output_video_link" binding:"required"`
	Model               ModelEntity `json:"model"`
}

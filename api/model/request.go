package model

type BaseRequest struct {
	InputFileName  string      `json:"input_file_name" binding:"required"`
	InputLink      string      `json:"input_link" binding:"required"`
	OutputFileName string      `json:"output_file_name" binding:"required"`
	OutputLink     string      `json:"output_link" binding:"required"`
	Model          ModelEntity `json:"model" binding:"required"`
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
}

type TTTRequest struct {
	BaseRequest
	BaseLang
}

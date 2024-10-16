package model

type ModelRequest struct {
	InputDir  string `json:"input_dir" binding:"required"`
	OutputDir string `json:"output_dir" binding:"required"`
}

type STTRequest struct {
	InputFile  string `json:"input_file" binding:"required"`
	OutputFile string `json:"output_file" binding:"required"`
}

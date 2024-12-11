package command

import (
	"context"
	"fmt"
	"mlvt-api/internal/python"
	"os/exec"
	"time"
)

// ExecuteCommand runs a Python script with the given arguments.
func ExecuteCommand(pythonVersion python.PythonVersion, scriptPath string, args []string, timeout time.Duration) error {
	pythonExec := python.GetPythonExecutable(pythonVersion)

	// Prepare the command context with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmdArgs := append([]string{scriptPath}, args...)

	fmt.Printf("- - -cmd: %v \n", cmdArgs)

	cmd := exec.CommandContext(ctx, pythonExec, cmdArgs...)

	fmt.Printf("- - - cmd: %v \n", cmd)

	// Execute the command.
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v, output: %s", err, string(output))
	}

	return nil
}

// RunSTT executes the stt.py script.
func RunSTT(pythonVersion python.PythonVersion, scriptPath, inputFile, outputFile string) error {
	fmt.Printf("- - debug: %v; %v \n", inputFile, outputFile)
	args := []string{inputFile, outputFile}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

// RunTTS executes the tts.py script.
func RunTTS(pythonVersion python.PythonVersion, scriptPath, inputFile, outputFile string) error {
	args := []string{inputFile, outputFile}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

// RunTTT executes the ttt.py script.
func RunTTT(pythonVersion python.PythonVersion, scriptPath, inputFile, outputFile, sourceLang, targetLang string) error {
	args := []string{inputFile, outputFile, sourceLang, targetLang}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

// RunLS executes the ls.py (lip-sync) script.
func RunLS(pythonVersion python.PythonVersion, scriptPath, inputVideoFile, inputAudioFile, outputFile string) error {
	fmt.Printf("- - debug (LS): %v; %v; %v\n", inputVideoFile, inputAudioFile, outputFile)
	args := []string{inputVideoFile, inputAudioFile, outputFile}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

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

	// Prepare the command arguments.
	cmdArgs := append([]string{scriptPath}, args...)

	fmt.Printf("- - -cmd: %v \n", cmdArgs)

	// Create the command.
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
func RunSTT(pythonVersion python.PythonVersion, inputFile, outputFile string) error {
	scriptPath := "/home/ec2-user/mlvt-api/scripts/stt/stt.py"
	fmt.Printf("- - debug: %v; %v \n", inputFile, outputFile)
	args := []string{inputFile, outputFile}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

// RunTTS executes the tts.py script.
func RunTTS(pythonVersion python.PythonVersion, inputFile, outputFile string) error {
	scriptPath := "./scripts/tts/tts.py"
	args := []string{inputFile, outputFile}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

// RunTTT executes the ttt.py script.
func RunTTT(pythonVersion python.PythonVersion, inputFile, outputFile, sourceLang, targetLang string) error {
	scriptPath := "./scripts/ttt/ttt.py"
	args := []string{inputFile, outputFile, sourceLang, targetLang}
	timeout := 5 * time.Minute // Adjust as needed.

	return ExecuteCommand(pythonVersion, scriptPath, args, timeout)
}

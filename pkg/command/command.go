package command

import "os/exec"

// Define version python with limited options
type PythonVersion string

const (
	Py3_11 PythonVersion = "python3.11"
	Py3_10 PythonVersion = "python3.10"
	Py3_8  PythonVersion = "python3.9"
)

func RunCommand(py_version PythonVersion, filename, inputFile, outputFile string) *exec.Cmd {
	return exec.Command(string(py_version), filename, "--input_file", inputFile, "--output_file", outputFile)
}

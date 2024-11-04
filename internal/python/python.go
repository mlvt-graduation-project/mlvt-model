package python

// PythonVersion defines supported Python versions.
type PythonVersion string

const (
	Py3_12 PythonVersion = "python3.12"
	Py3_11 PythonVersion = "python3.11"
	Py3_10 PythonVersion = "python3.10"
	Py3_9  PythonVersion = "python3.9"
	Py3    PythonVersion = "python3"
)

// GetPythonExecutable returns the Python executable based on the version.
func GetPythonExecutable(version PythonVersion) string {
	return string(version)
}

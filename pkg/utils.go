package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// UploadFile uploads a file to a presigned URL.
func UploadFile(filePath, url string) error {
	// Read the file.
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create a PUT request.
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to upload file, status code: %d", resp.StatusCode)
	}

	return nil
}

// DownloadFile downloads a file from a presigned URL.
func DownloadFile(url, destPath string) error {
	// Get the data.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file.
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file.
	_, err = io.Copy(out, resp.Body)
	return err
}

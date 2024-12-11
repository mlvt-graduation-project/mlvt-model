package queue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"mlvt-api/api/model"
)

// CallbackPayload represents the payload sent to the callback API.
type CallbackPayload struct {
	JobID  string      `json:"job_id"`
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"` // Optional: include result
}

// CallCallbackAPI sends the job status to the callback URL.
func (p *Processor) CallCallbackAPI(callbackURL string, job *model.Job) error {
	payload := CallbackPayload{
		JobID:  job.ID,
		Status: job.Status,
		Error:  job.Error,
		Result: job.Result, // Include the result if available
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal callback payload: %v", err)
	}

	resp, err := http.Post(callbackURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to call callback API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("callback API returned status: %s", resp.Status)
	}

	return nil
}

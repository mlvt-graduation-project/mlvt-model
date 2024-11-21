package model

import "errors"

// Define specific errors
var (
	ErrInvalidRequestType = errors.New("invalid request type")
	ErrUnknownJobType     = errors.New("unknown job type")
)

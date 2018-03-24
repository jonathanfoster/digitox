package model

import "net/http"

// Error represents an API error response.
type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// NewError creates an Error instance.
func NewError(statusCode int) *Error {
	return &Error{
		Message:    http.StatusText(statusCode),
		StatusCode: statusCode,
	}
}

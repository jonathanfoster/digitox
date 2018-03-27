package models

// DefaultStatus is the current API status.
var DefaultStatus *Status

// Status represents the API status.
type Status struct {
	Version string `json:"version"`
}

// NewStatus creates a Status instance.
func NewStatus() *Status {
	return &Status{}
}

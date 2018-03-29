package status

// Current is the current API status.
var Current *Status

// Status represents the API status.
type Status struct {
	Version string `json:"version"`
}

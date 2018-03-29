package status

// Current is the current server status.
var Current *Status

// Status represents the server status.
type Status struct {
	Version string `json:"version"`
}

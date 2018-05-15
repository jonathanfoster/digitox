package handlers

import (
	"net/http"
)

// ServerVersion is the API server version.
var ServerVersion string

// VersionPayload is the payload returned from the Version handler.
type VersionPayload struct {
	Version string `json:"version"`
}

// Version handles the GET / route.
func Version(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, &VersionPayload{
		Version: ServerVersion,
	})
}

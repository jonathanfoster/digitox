package handlers

import (
	"net/http"
)

// ServerVersion is the API server version.
var ServerVersion string

// Version handles the GET / route.
func Version(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, struct {
		Version string
	}{
		Version: ServerVersion,
	})
}

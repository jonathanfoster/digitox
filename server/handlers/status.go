package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/models/status"
)

// Status handles the GET / route.
func Status(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, status.Current)
}

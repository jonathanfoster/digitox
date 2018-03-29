package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/server/status"
)

// Status handles the GET / route.
func Status(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, status.Current)
}

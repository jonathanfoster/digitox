package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/models"
)

// Status handles the GET / route.
func Status(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, models.DefaultStatus)
}

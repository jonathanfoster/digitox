package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/api/httputil"
	"github.com/jonathanfoster/freedom/models"
)

// Status handles the GET / route.
func Status(w http.ResponseWriter, r *http.Request) {
	httputil.JSON(w, http.StatusOK, models.DefaultStatus)
}

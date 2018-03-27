package httputil

import (
	"net/http"

	"github.com/jonathanfoster/freedom/models"
)

// Error writes application/json error to response writer.
func Error(w http.ResponseWriter, statusCode int) {
	JSON(w, statusCode, models.NewError(statusCode))
}

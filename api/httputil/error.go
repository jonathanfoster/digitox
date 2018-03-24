package httputil

import (
	"net/http"

	"github.com/jonathanfoster/freedom/model"
)

// Error writes application/json error to response writer.
func Error(w http.ResponseWriter, statusCode int) {
	JSON(w, statusCode, model.NewError(statusCode))
}

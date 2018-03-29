package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Error writes application/json error to response writer.
func Error(w http.ResponseWriter, statusCode int) {
	JSON(w, statusCode, &errorResponse{
		Message:    http.StatusText(statusCode),
		StatusCode: statusCode,
	})
}

// JSON writes application/json to response writer.
func JSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Error("error json encoding: ", err.Error())
	}
}

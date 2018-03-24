package httputil

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// JSON writes application/json to response writer.
func JSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Error("error json encoding: ", err.Error())
	}
}

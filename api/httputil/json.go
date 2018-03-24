package httputil

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// JSON writes application/json to writer.
func JSON(w http.ResponseWriter, statusCode int, i interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Error("error json encoding: ", err.Error())
	}
}

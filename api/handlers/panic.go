package handlers

import (
	"errors"
	"net/http"
)

// Panic handles the GET /panic route.
func Panic(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("test"))
}

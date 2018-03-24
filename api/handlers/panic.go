package handlers

import (
	"errors"
	"net/http"
)

// Error handles the GET /error route.
func Panic(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("test"))
}

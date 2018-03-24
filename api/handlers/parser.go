package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

// ParseID parses ID from route variables.
func ParseID(r *http.Request) (uuid.UUID, bool) {
	rv := mux.Vars(r)["id"]
	id, err := uuid.FromString(rv)
	return id, err == nil
}

// ParseName parses name from route variables.
func ParseName(r *http.Request) (string, bool) {
	rv := mux.Vars(r)["name"]
	return rv, rv != ""
}

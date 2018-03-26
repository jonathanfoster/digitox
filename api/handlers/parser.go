package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ParseID parses ID from route variables.
func ParseID(r *http.Request) (string, bool) {
	rv := mux.Vars(r)["id"]
	return rv, rv != ""
}

// ParseName parses name from route variables.
func ParseName(r *http.Request) (string, bool) {
	rv := mux.Vars(r)["name"]
	return rv, rv != ""
}

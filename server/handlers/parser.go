package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// ParseID parses ID from route variables.
func ParseID(r *http.Request) (uuid.UUID, error) {
	rv := mux.Vars(r)["id"]

	id, err := uuid.FromString(rv)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error parsing id")
	}

	return id, nil
}

// ParseName parses name from route variables.
func ParseName(r *http.Request) (string, bool) {
	rv := mux.Vars(r)["name"]

	return rv, rv != ""
}

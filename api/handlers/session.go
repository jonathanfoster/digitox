package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/api/httputil"
)

// ListSessions handles the GET /sessions route.
func ListSessions(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// FindSession handles the GET /session/{id} route.
func FindSession(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// CreateSession handles the POST /sessions/{id} route.
func CreateSession(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// DeleteSession handles the DELETE /sessions/{id} route.
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// UpdateSession handles the PUT /sessions/{id} route.
func UpdateSession(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}
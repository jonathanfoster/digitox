package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/api/httputil"
)

// ListBlocklists handles the GET /blocklists route.
func ListBlocklists(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// FindBlocklist handles the GET /blocklists/{id} route.
func FindBlocklist(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// CreateBlocklist handles the POST /blocklists/{id} route.
func CreateBlocklist(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// DeleteBlocklist handles the DELETE /blocklists/{id} route.
func DeleteBlocklist(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// UpdateBlocklist handles the PUT /blocklists/{id} route.
func UpdateBlocklist(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

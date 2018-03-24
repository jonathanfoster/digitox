package handlers

import (
	"net/http"

	"github.com/jonathanfoster/freedom/api/httputil"
)

// ListDevices handles the GET /devices route.
func ListDevices(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// FindDevice handles the GET /devices/{id} route.
func FindDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// CreateDevice handles the POST /devices/{id} route.
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// DeleteDevice handles the DELETE /devices/{id} route.
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// UpdateDevice handles the PUT /devices/{id} route.
func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

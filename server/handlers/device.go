package handlers

import (
	"net/http"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/store"
)

// ListDevices handles the GET /devices route.
func ListDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := device.All()
	if err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warn("devices do not exist: ", err.Error())
			JSON(w, http.StatusOK, []*blocklist.Blocklist{})
			return
		}

		log.Error("error listing devices: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if devices == nil {
		devices = []*device.Device{}
	}

	JSON(w, http.StatusOK, devices)
}

// FindDevice handles the GET /devices/{id} route.
func FindDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	Error(w, http.StatusNotImplemented)
}

// CreateDevice handles the POST /devices/{id} route.
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusNotImplemented)
}

// DeleteDevice handles the DELETE /devices/{id} route.
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	Error(w, http.StatusNotImplemented)
}

// UpdateDevice handles the PUT /devices/{id} route.
func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	Error(w, http.StatusNotImplemented)
}

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/store"
)

// ListDevices handles the GET /devices/ route.
func ListDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := device.All()
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
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
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	dev, err := device.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("device %s not found: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding device %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, dev)
}

// CreateDevice handles the POST /devices/ route.
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading dev body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	var dev device.Device
	if err := json.Unmarshal(buf, &dev); err != nil {
		log.Warn("error unmarshaling dev body: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if valid, err := dev.Validate(); !valid {
		log.Warn("dev not valid: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if err := dev.Save(); err != nil {
		log.Error("error saving dev: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusCreated, dev)
}

// DeleteDevice handles the DELETE /devices/{id} route.
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	if err := device.Remove(id); err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("device %s not found: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error removing device %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateDevice handles the PUT /devices/{id} route.
func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	dev, err := device.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("device %s not found: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding device %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(buf, &dev); err != nil {
		log.Warn("error unmarshaling body: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if valid, err := dev.Validate(); !valid {
		log.Warn("device not valid: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if err := dev.Save(); err != nil {
		log.Error("error saving device: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, dev)
}

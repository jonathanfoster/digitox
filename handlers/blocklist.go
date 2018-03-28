package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models"
	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/store"
)

// ListBlocklists handles the GET /blocklists route.
func ListBlocklists(w http.ResponseWriter, r *http.Request) {
	lists, err := blocklist.All()
	if err != nil {
		log.Error("error listing blocklists: ", err.Error())
		Error(w, http.StatusInternalServerError)
	}

	JSON(w, http.StatusOK, lists)
}

// FindBlocklist handles the GET /blocklists/{id} route.
func FindBlocklist(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	list, err := blocklist.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("blocklist %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding blocklist %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, list)
}

// CreateBlocklist handles the POST /blocklists route.
func CreateBlocklist(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading blocklist body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	var list blocklist.Blocklist
	if err := json.Unmarshal(buf, &list); err != nil {
		log.Error("error unmarshaling blocklist body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := list.Save(); err != nil {
		if models.IsValidator(err) {
			log.Warn("blocklist not valid: ", err.Error())
			Error(w, http.StatusUnprocessableEntity)
			return
		}

		log.Error("error saving blocklist: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusCreated, list)
}

// DeleteBlocklist handles the DELETE /blocklists/{id} route.
func DeleteBlocklist(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	if err := blocklist.Remove(id); err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("blocklist %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error removing blocklist %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateBlocklist handles the PUT /blocklists/{id} route.
func UpdateBlocklist(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	list, err := blocklist.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("blocklist %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding blocklist %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(buf, &list); err != nil {
		log.Error("error unmarshaling body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := list.Save(); err != nil {
		if models.IsValidator(err) {
			log.Warn("blocklist not valid: ", err.Error())
			Error(w, http.StatusUnprocessableEntity)
			return
		}

		log.Error("error saving blocklist: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, list)
}

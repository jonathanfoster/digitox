package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/store"
)

// ListSessions handles the GET /sessions route.
func ListSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := session.All()
	if err != nil {
		log.Error("error listing sessions: ", err.Error())
		Error(w, http.StatusInternalServerError)
	}

	JSON(w, http.StatusOK, sessions)
}

// FindSession handles the GET /sessions/{id} route.
func FindSession(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	sess, err := session.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("session %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding session %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, sess)
}

// CreateSession handles the POST /sessions/{id} route.
func CreateSession(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusNotImplemented)
}

// DeleteSession handles the DELETE /sessions/{id} route.
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	if err := session.Remove(id); err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("session %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error removing session %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateSession handles the PUT /sessions/{id} route.
func UpdateSession(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	sess, err := session.Find(id)
	if err != nil {
		if errors.Cause(err) == store.ErrNotExist {
			log.Warnf("session %s does not exist: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding session %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(buf, &sess); err != nil {
		log.Error("error unmarshaling body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if err := sess.Save(); err != nil {
		if models.IsValidator(err) {
			log.Warn("session not valid: ", err.Error())
			Error(w, http.StatusUnprocessableEntity)
			return
		}

		log.Error("error saving session: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, sess)
}

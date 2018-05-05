package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/store"
)

// ListSessions handles the GET /sessions/ route.
func ListSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := session.All()
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			log.Warn("sessions not found: ", err.Error())
			JSON(w, http.StatusOK, []*session.Session{})
			return
		}

		log.Error("error listing sessions: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if sessions == nil {
		sessions = []*session.Session{}
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
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("session %s not found: %s", id, err.Error())
			Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding session %s: %s", id, err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, sess)
}

// CreateSession handles the POST /sessions/ route.
func CreateSession(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("error reading session body: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	var sess session.Session
	if err := json.Unmarshal(buf, &sess); err != nil {
		log.Warn("error unmarshaling session body: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	sess.ID = uuid.NewV4()

	if valid, err := sess.Validate(); !valid {
		log.Warn("session not valid: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if err := sess.Save(); err != nil {
		log.Error("error saving session: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusCreated, sess)
}

// RemoveSession handles the DELETE /sessions/{id} route.
func RemoveSession(w http.ResponseWriter, r *http.Request) {
	id, ok := ParseID(r)
	if !ok {
		Error(w, http.StatusBadRequest)
		return
	}

	if err := session.Remove(id); err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("session %s not found: %s", id, err.Error())
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
		if errors.Cause(err) == store.ErrNotFound {
			log.Warnf("session %s not found: %s", id, err.Error())
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
		log.Warn("error unmarshaling body: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if valid, err := sess.Validate(); !valid {
		log.Warn("session not valid: ", err.Error())
		Error(w, http.StatusUnprocessableEntity)
		return
	}

	if err := sess.Save(); err != nil {
		log.Error("error saving session: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, sess)
}

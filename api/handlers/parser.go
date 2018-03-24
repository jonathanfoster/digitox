package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// ParseID parses ID from the route variables for the current request.
func ParseID(r *http.Request) (uuid.UUID, bool) {
	idVar := mux.Vars(r)["id"]
	id, err := uuid.FromString(idVar)
	if err != nil {
		log.Warnf("id not a valid uuid: ", idVar)
	}
	return id, err != nil
}

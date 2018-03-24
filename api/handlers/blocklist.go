package handlers

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/api/httputil"
	"github.com/jonathanfoster/freedom/model/blocklist"
)

// ListBlocklists handles the GET /blocklists route.
func ListBlocklists(w http.ResponseWriter, r *http.Request) {
	lists, err := blocklist.All()
	if err != nil {
		log.Error("error listing blocklists: ", err.Error())
		httputil.Error(w, http.StatusInternalServerError)
	}

	httputil.JSON(w, http.StatusOK, lists)
}

// FindBlocklist handles the GET /blocklists/{name} route.
func FindBlocklist(w http.ResponseWriter, r *http.Request) {
	name, ok := ParseName(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	list, err := blocklist.Find(name)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warnf("blocklist does not exist: %s: %s", name, err.Error())
			httputil.Error(w, http.StatusNotFound)
			return
		}

		log.Errorf("error finding blocklist %s: %s", name, err.Error())
		httputil.Error(w, http.StatusInternalServerError)
		return
	}

	httputil.JSON(w, http.StatusOK, list)
}

// CreateBlocklist handles the POST /blocklists/{name} route.
func CreateBlocklist(w http.ResponseWriter, r *http.Request) {
	httputil.Error(w, http.StatusNotImplemented)
}

// DeleteBlocklist handles the DELETE /blocklists/{name} route.
func DeleteBlocklist(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseName(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

// UpdateBlocklist handles the PUT /blocklists/{name} route.
func UpdateBlocklist(w http.ResponseWriter, r *http.Request) {
	_, ok := ParseName(r)
	if !ok {
		httputil.Error(w, http.StatusBadRequest)
		return
	}

	httputil.Error(w, http.StatusNotImplemented)
}

package handlers

import (
	"net/http"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/proxy"
)

// ProxyActive handles the GET /proxy/session route.
func ProxyActive(w http.ResponseWriter, r *http.Request) {
	list, err := proxy.ProxyController.ActiveBlocklist()
	if err != nil {
		log.Error("error listing blocklists: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if list == nil {
		list = []string{}
	}

	JSON(w, http.StatusOK, list)
}

// ProxyBlocklist handles the GET /proxy/blocklist route.
func ProxyBlocklist(w http.ResponseWriter, r *http.Request) {
	list, err := proxy.ProxyController.ReadBlocklistFile()
	if err != nil {
		log.Error("error reading blocklist file: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	if list == nil {
		list = []string{}
	}

	JSON(w, http.StatusOK, list)
}

// ReconfigureProxy handles the POST /proxy/reconfigure route.
func ReconfigureProxy(w http.ResponseWriter, r *http.Request) {
	// Reload Squid proxy config by calling `/usr/sbin/squid -k reconfigure`
	// This operation will have negative  side effects like ports closing
	// and loss of information on in-transit requests (https://wiki.squid-cache.org/Features/HotConf)
	cmd := exec.Command("/usr/sbin/squid", "-k", "reconfigure")
	if err := cmd.Run(); err != nil {
		log.Error("error reconfiguring proxy: ", err.Error())
		Error(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

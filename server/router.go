package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jonathanfoster/digitox/server/handlers"
)

// Router represents a router with preconfigured routes.
type Router struct {
	router *mux.Router
}

// NewRouter creates a Router instance.
func NewRouter() *Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Version).Methods("GET")
	router.HandleFunc("/oauth/token", handlers.Token).Methods("GET", "POST")

	router.HandleFunc("/blocklists/", handlers.ListBlocklists).Methods("GET")
	router.HandleFunc("/blocklists/{id}", handlers.FindBlocklist).Methods("GET")
	router.HandleFunc("/blocklists/", handlers.CreateBlocklist).Methods("POST")
	router.HandleFunc("/blocklists/{id}", handlers.UpdateBlocklist).Methods("PUT")
	router.HandleFunc("/blocklists/{id}", handlers.RemoveBlocklist).Methods("DELETE")

	router.HandleFunc("/devices/", handlers.ListDevices).Methods("GET")
	router.HandleFunc("/devices/{name}", handlers.FindDevice).Methods("GET")
	router.HandleFunc("/devices/", handlers.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{name}", handlers.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{name}", handlers.RemoveDevice).Methods("DELETE")

	router.HandleFunc("/proxy/active", handlers.ProxyActive).Methods("GET")
	router.HandleFunc("/proxy/blocklist", handlers.ProxyBlocklist).Methods("GET")
	router.HandleFunc("/proxy/reconfigure", handlers.ReconfigureProxy).Methods("POST")

	router.HandleFunc("/sessions/", handlers.ListSessions).Methods("GET")
	router.HandleFunc("/sessions/{id}", handlers.FindSession).Methods("GET")
	router.HandleFunc("/sessions/", handlers.CreateSession).Methods("POST")
	router.HandleFunc("/sessions/{id}", handlers.UpdateSession).Methods("PUT")
	router.HandleFunc("/sessions/{id}", handlers.RemoveSession).Methods("DELETE")

	return &Router{
		router,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

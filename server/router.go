package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jonathanfoster/freedom/server/handlers"
)

// Router represents a gorilla/mux router with preconfigured routes.
type Router struct {
	router *mux.Router
}

// NewRouter creates a Router instance.
func NewRouter() *Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Status).Methods("GET")
	router.HandleFunc("/panic", handlers.Panic).Methods("GET")

	router.HandleFunc("/blocklists", handlers.ListBlocklists).Methods("GET")
	router.HandleFunc("/blocklists/{id}", handlers.FindBlocklist).Methods("GET")
	router.HandleFunc("/blocklists", handlers.CreateBlocklist).Methods("POST")
	router.HandleFunc("/blocklists/{id}", handlers.UpdateBlocklist).Methods("PUT")
	router.HandleFunc("/blocklists/{id}", handlers.DeleteBlocklist).Methods("DELETE")

	router.HandleFunc("/devices", handlers.ListDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", handlers.FindDevice).Methods("GET")
	router.HandleFunc("/devices", handlers.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", handlers.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{id}", handlers.DeleteDevice).Methods("DELETE")

	router.HandleFunc("/sessions", handlers.ListSessions).Methods("GET")
	router.HandleFunc("/sessions/{id}", handlers.FindSession).Methods("GET")
	router.HandleFunc("/sessions", handlers.CreateSession).Methods("POST")
	router.HandleFunc("/sessions/{id}", handlers.UpdateSession).Methods("PUT")
	router.HandleFunc("/sessions/{id}", handlers.DeleteSession).Methods("DELETE")

	return &Router{
		router,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jonathanfoster/freedom/api/handlers"
)

// Router represents a gorilla/mux router with preconfigured routes.
type Router struct {
	router *mux.Router
}

// NewRouter creates a Router instance.
func NewRouter() *Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Status)
	router.HandleFunc("/panic", handlers.Panic)

	router.HandleFunc("/blocklists", handlers.ListBlocklists)
	router.HandleFunc("/blocklists/{name}", handlers.FindBlocklist)
	router.HandleFunc("/blocklists", handlers.CreateBlocklist).Methods("POST")
	router.HandleFunc("/blocklists/{name}", handlers.UpdateBlocklist).Methods("PUT")
	router.HandleFunc("/blocklists/{name}", handlers.DeleteBlocklist).Methods("DELETE")

	router.HandleFunc("/devices", handlers.ListDevices)
	router.HandleFunc("/devices/{id}", handlers.FindDevice)
	router.HandleFunc("/devices", handlers.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", handlers.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{id}", handlers.DeleteDevice).Methods("DELETE")

	router.HandleFunc("/sessions", handlers.ListSessions)
	router.HandleFunc("/sessions/{id}", handlers.FindSession)
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

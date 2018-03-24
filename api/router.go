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
	router.HandleFunc("/blocklists/{id}", handlers.FindBlocklist)
	router.HandleFunc("/blocklists", handlers.CreateBlocklist).Methods("POST")
	router.HandleFunc("/blocklists/{id}", handlers.UpdateBlocklist).Methods("PUT")
	router.HandleFunc("/blocklists/{id}", handlers.DeleteBlocklist).Methods("DELETE")

	router.HandleFunc("/session", handlers.ListSessions)
	router.HandleFunc("/session/{id}", handlers.FindSession)
	router.HandleFunc("/session", handlers.CreateSession).Methods("POST")
	router.HandleFunc("/session/{id}", handlers.UpdateSession).Methods("PUT")
	router.HandleFunc("/session/{id}", handlers.DeleteSession).Methods("DELETE")

	return &Router{
		router,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

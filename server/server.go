package server

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/jonathanfoster/freedom/middleware"
)

// Server represents a server.
type Server struct{}

// New creates a Server instance.
func New() *Server {
	return &Server{}
}

// Run listens on the TCP network address addr and then
// calls Serve to handle requests on incoming connections.
func (s *Server) Run(addr string) {
	router := NewRouter()

	n := negroni.New()
	n.Use(middleware.NewRecovery())
	n.Use(middleware.NewLogger())
	n.UseHandler(router)

	srv := &http.Server{
		Addr:    addr,
		Handler: n,
	}

	log.Info("server listening on ", addr)
	if err := gracehttp.Serve(srv); err != nil {
		log.Fatal("error starting server: ", err.Error())
	}
}

package server

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/pkg/errors"
	"github.com/urfave/negroni"

	"github.com/jonathanfoster/freedom/server/middleware"
)

// Server represents a server.
type Server struct{}

// New creates a Server instance.
func New() *Server {
	return &Server{}
}

// Run listens on the TCP network address addr and then
// calls Serve to handle requests on incoming connections.
func (s *Server) Run(addr string) error {
	router := NewRouter()

	n := negroni.New()
	n.Use(middleware.NewRecovery())
	n.Use(middleware.NewLogger())
	n.UseHandler(router)

	srv := &http.Server{
		Addr:    addr,
		Handler: n,
	}

	if err := gracehttp.Serve(srv); err != nil {
		return errors.Wrap(err, "error starting server")
	}

	return nil
}

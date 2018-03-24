package api

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/jonathanfoster/freedom/api/middleware"
)

// Server represents a server.
type Server struct{}

// NewServer creates a Server instance.
func NewServer() *Server {
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
	gracehttp.Serve(srv)
}

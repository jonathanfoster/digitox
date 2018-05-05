package server

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/pkg/errors"
	"github.com/urfave/negroni"

	"github.com/jonathanfoster/digitox/server/middleware"
	"github.com/jonathanfoster/digitox/server/oauth"
)

// Server represents a server.
type Server struct {
	config *Config
}

// New creates a Server instance.
func New(config *Config) *Server {
	return &Server{config: config}
}

// Run listens on the TCP network address addr and then
// calls Serve to handle requests on incoming connections.
func (s *Server) Run() error {
	oauth.InitOAuthServer(s.config.TokenSigningKey, s.config.ClientID, s.config.ClientSecret)

	router := NewRouter()

	n := negroni.New()
	n.Use(middleware.NewAuth(s.config.TokenVerifyingKey))
	n.Use(middleware.NewRecovery())
	n.Use(middleware.NewLogger())
	n.UseHandler(router)

	srv := &http.Server{
		Addr:    s.config.Addr,
		Handler: n,
	}

	if err := gracehttp.Serve(srv); err != nil {
		return errors.Wrap(err, "error starting server")
	}

	return nil
}

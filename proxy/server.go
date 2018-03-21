package proxy

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

// Server represents a proxy server.
type Server struct {
	Server *http.Server
}

// NewServer creates a new instance of Server.
func NewServer() (*Server, error) {
	fwd := NewForwarder()
	blocker := NewBlocker(fwd)

	srv := &http.Server{
		Handler: blocker,
	}

	return &Server{
		Server: srv,
	}, nil
}

// Start listens on the TCP network address addr and then calls Serve to handle requests on incoming connections.
func (srv *Server) Start(addr string) {
	srv.Server.Addr = addr

	go func() {
		log.Info("server listening on ", srv.Server.Addr)
		if err := srv.Server.ListenAndServe(); err != nil {
			// ListentAndServe always returns ErrrServerClosed when calling shutdown.
			if err != http.ErrServerClosed {
				log.WithError(err).Fatal("error starting server: ", err)
			}
		}
	}()

	// Gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("server shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Server.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("error shutting down server: ", err)
	}
}
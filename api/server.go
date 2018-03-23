package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	srv := &http.Server{
		Addr:    addr,
		Handler: n,
	}

	go func() {
		log.Info("server listening on ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("error starting server: ", err)
		}
	}()

	// Gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("error shutting down server: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

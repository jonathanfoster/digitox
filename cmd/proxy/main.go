package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/jonathanfoster/freedom/proxy"
	log "github.com/sirupsen/logrus"
)

var (
	version string
	app     = kingpin.New("freedom-proxy", "Freedom proxy enforces sessions by denying access to hosts found in block lists.").Version(version)
	port    = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug log messages enabled")
	}

	srv, err := proxy.NewServer()
	if err != nil {
		log.WithError(err).Fatal("error creating proxy server: ", err.Error())
	}

	addr := ":" + *port
	srv.Start(addr)
}

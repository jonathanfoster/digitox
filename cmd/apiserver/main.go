package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/server"
	"github.com/jonathanfoster/freedom/server/status"
)

var (
	version string
	app     = kingpin.New("freedom-apiserver", "Freedom API server provides a REST API for managing Freedom proxy.").Version(version)
	port    = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log messages enabled")
	}

	status.Current = &status.Status{
		Version: version,
	}

	srv := server.New()
	srv.Run(":" + *port)
}

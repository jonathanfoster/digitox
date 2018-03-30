package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/server"
	"github.com/jonathanfoster/freedom/server/status"
	"github.com/jonathanfoster/freedom/store"
)

var (
	version  string
	app      = kingpin.New("freedom-apiserver", "Freedom API server provides a REST API for managing Freedom proxy.").Version(version)
	port     = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose  = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
	database = app.Flag("database", "Database source path").Short('d').Default("/etc/freedom/freedom.db").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log messages enabled")
	}

	db, err := gorm.Open("sqlite3", *database)
	if err != nil {
		log.Fatalf("error connecting to database %s: %s", *database, err.Error())
	}

	defer db.Close()

	db.AutoMigrate(&blocklist.Blocklist{}, &session.Session{})

	store.DB = db

	status.Current = &status.Status{
		Version: version,
	}

	srv := server.New()
	srv.Run(":" + *port)
}

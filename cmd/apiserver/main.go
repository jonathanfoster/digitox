package main

import (
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/proxy"
	"github.com/jonathanfoster/digitox/server"
	"github.com/jonathanfoster/digitox/server/status"
	"github.com/jonathanfoster/digitox/store"
)

var (
	version    string
	app        = kingpin.New("digitox-apiserver", "Digitox API server provides a REST API for managing resources.").Version(version)
	port       = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose    = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
	sessions   = app.Flag("sessions", "Sessions store directory.").Default("/etc/digitox/sessions/").String()
	blocklists = app.Flag("blocklists", "Blocklists store directory.").Default("/etc/digitox/blocklists/").String()
	active     = app.Flag("active", "Active blocklist file name.").Default("/etc/digitox/active").String()
	devices    = app.Flag("devices", "Devices password file name.").Default("/etc/digitox/passwd").String()
	tick       = app.Flag("tick", "Tick duration of blocklist update ticker.").Short('t').Default("1s").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log messages enabled")
	}

	log.Infof("initializing blocklist store %s", *blocklists)
	store.Blocklist = store.NewFileStore(*blocklists)
	if err := store.Blocklist.Init(); err != nil {
		log.Fatal("error initializing blocklist store: ", err.Error())
	}

	log.Infof("initializing session store %s", *sessions)
	store.Session = store.NewFileStore(*sessions)
	if err := store.Session.Init(); err != nil {
		log.Fatal("error initializing session store: ", err.Error())
	}

	log.Infof("initializing device store %s", *devices)
	store.Device = store.NewHtpasswdStore(*devices)
	if err := store.Device.Init(); err != nil {
		log.Fatal("error initializing device store: ", err.Error())
	}

	d, err := time.ParseDuration(*tick)
	if err != nil {
		log.Warnf("error parsing duration %s: using default value 1s: %s", err.Error())
		d = time.Second * 1
	}

	log.Info("starting proxy controller")
	ctrl := proxy.NewController(*active)
	ctrl.Tick = d
	ctrl.Run()

	status.Current = &status.Status{
		Version: version,
	}

	addr := ":" + *port
	code := 0

	log.Info("server listening on ", addr)
	srv := server.New()
	if err := srv.Run(addr); err != nil {
		log.Error("error starting server: ", err)
		code = 1
	}

	log.Info("stopping proxy controller")
	if err := ctrl.Stop(); err != nil {
		log.Error("error stopping proxy controller: ", err)
		code = 1
	}

	os.Exit(code)
}

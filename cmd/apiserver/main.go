package main

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/proxy"
	"github.com/jonathanfoster/digitox/server"
	"github.com/jonathanfoster/digitox/server/oauth"
	"github.com/jonathanfoster/digitox/server/status"
	"github.com/jonathanfoster/digitox/store"
)

var (
	version          string
	app              = kingpin.New("digitox-apiserver", "Digitox API server provides a REST API for managing resources.").Version(version)
	port             = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose          = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
	sessions         = app.Flag("sessions", "Sessions store directory.").Default("/etc/digitox/sessions/").String()
	blocklists       = app.Flag("blocklists", "Blocklists store directory.").Default("/etc/digitox/blocklists/").String()
	active           = app.Flag("active", "Active blocklist file name.").Default("/etc/digitox/active").String()
	devices          = app.Flag("devices", "Devices password file name.").Default("/etc/digitox/passwd").String()
	tick             = app.Flag("tick", "Tick duration of blocklist update ticker.").Short('t').Default("1s").String()
	signingKeyPath   = app.Flag("signing-key", "RSA private key path for signing JWT tokens.").Default("/etc/digitox/signing-key.pem").String()
	verifyingKeyPath = app.Flag("verifying-key", "RSA public key path verifying JWT tokens.").Default("/etc/digitox/verifying-key.pem").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log messages enabled")
	}

	initStores(blocklists, sessions, devices)

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

	config := server.NewConfig()

	config.TokenSigningKey = initSigningKey(signingKeyPath)
	config.TokenVerifyingKey = initVerifyingKey(verifyingKeyPath)

	addr := ":" + *port
	code := 0
	config.Addr = ":" + *port

	log.Info("server listening on ", addr)
	srv := server.New(config)
	if err := srv.Run(); err != nil {
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

func initStores(blocklists *string, sessions *string, devices *string) {
	log.Infof("initializing blocklist store %s", *blocklists)
	store.Blocklist = store.NewFileStore(*blocklists)
	if err := store.Blocklist.Init(); err != nil {
		log.Error("error initializing blocklist store: ", err.Error())
	}

	log.Infof("initializing session store %s", *sessions)
	store.Session = store.NewFileStore(*sessions)
	if err := store.Session.Init(); err != nil {
		log.Error("error initializing session store: ", err.Error())
	}

	log.Infof("initializing device store %s", *devices)
	store.Device = store.NewHtpasswdStore(*devices)
	if err := store.Device.Init(); err != nil {
		log.Error("error initializing device store: ", err.Error())
	}
}

func initSigningKey(signingKeyPath *string) (signingKey *rsa.PrivateKey) {
	if signingKeyPath != nil && *signingKeyPath != "" {
		signingKeyBytes, err := ioutil.ReadFile(*signingKeyPath)
		if err != nil {
			log.Error("error reading signing key file: ", err.Error())
		}

		if len(signingKeyBytes) > 0 {
			signingKey, err = jwt.ParseRSAPrivateKeyFromPEM(signingKeyBytes)
			if err != nil {
				log.Error("error parsing RSA private key from signing key bytes: ", err.Error())
			}
		}
	}

	if signingKey == nil {
		log.Warnf("signing key not provided: using default signing key")
		signingKey = oauth.DefaultSigningKey
	}

	return
}

func initVerifyingKey(verifyingKeyPath *string) (verifyingKey *rsa.PublicKey) {
	if verifyingKeyPath != nil && *verifyingKeyPath != "" {
		verifyingKeyBytes, err := ioutil.ReadFile(*verifyingKeyPath)
		if err != nil {
			log.Error("error reading verifying key file: ", err.Error())
		}

		if len(verifyingKeyBytes) > 0 {
			verifyingKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyingKeyBytes)
			if err != nil {
				log.Error("error parsing RSA public key from verifying key bytes: ", err.Error())
			}
		}
	}

	if verifyingKey == nil {
		log.Warnf("verifying key not provided: using default verifying key")
		verifyingKey = oauth.DefaultVerifyingKey
	}

	return
}

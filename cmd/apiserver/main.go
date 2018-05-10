package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	tickerDuration   = app.Flag("ticker-duration", "Duration of blocklist update ticker.").Short('t').Default("1s").String()
	signingKeyPath   = app.Flag("signing-key", "RSA private key path for signing JWT tokens.").Default("/etc/digitox/signing-key.pem").String()
	verifyingKeyPath = app.Flag("verifying-key", "RSA public key path verifying JWT tokens.").Default("/etc/digitox/verifying-key.pem").String()
	clientID         = app.Flag("client-id", "OAuth client ID.").String()
	clientSecret     = app.Flag("client-secret", "OAuth client secret.").String()
	dataSource       = app.Flag("data-source", "Database data source name.").Default("/etc/digitox/sessions.db").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	status.Current = &status.Status{
		Version: version,
	}

	config := server.NewConfig()

	config.Addr = ":" + *port
	config.DataSource = *dataSource
	config.Verbose = *verbose

	if config.Verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log messages enabled")
	}

	initCredentials(config, clientID, clientSecret)
	initStores(blocklists, sessions, devices)
	initTickerDuration(config, tickerDuration)
	initTokenSigningKey(config, signingKeyPath)
	initTokenVerifyingKey(config, verifyingKeyPath)

	log.Info("initializing database connection")
	if err := server.InitDB(config.DataSource, config.Verbose); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("starting proxy controller")
	ctrl := proxy.NewController(*active)
	ctrl.TickerDuration = config.TickerDuration
	ctrl.Run()

	exitCode := 0

	log.Info("server listening on ", config.Addr)
	srv := server.New(config)
	if err := srv.Run(); err != nil {
		log.Error("error starting server: ", err.Error())
		exitCode = 1
	}

	log.Info("stopping proxy controller")
	if err := ctrl.Stop(); err != nil {
		log.Error("error stopping proxy controller: ", err.Error())
		exitCode = 1
	}

	os.Exit(exitCode)
}

func initCredentials(config *server.Config, clientID *string, clientSecret *string) {
	if clientID == nil || *clientID == "" {
		log.Warnf("client ID not provided: using default client ID %s", oauth.DefaultClientID)
		config.ClientID = oauth.DefaultClientID
	} else {
		config.ClientID = *clientID
	}

	if clientSecret == nil || *clientSecret == "" {
		log.Warnf("client secret not provided: using default client secret %s", oauth.DefaultClientSecret)
		config.ClientSecret = oauth.DefaultClientSecret
	} else {
		config.ClientSecret = *clientSecret
	}
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

func initTickerDuration(config *server.Config, tick *string) {
	d, err := time.ParseDuration(*tick)
	if err != nil {
		log.Warnf("error parsing duration %s: using default value 1s: %s", err.Error())
		d = time.Second * 1
	}

	config.TickerDuration = d
}

func initTokenSigningKey(config *server.Config, signingKeyPath *string) {
	if signingKeyPath != nil && *signingKeyPath != "" {
		signingKeyBytes, err := ioutil.ReadFile(*signingKeyPath)
		if err != nil {
			log.Warn("error reading signing key file: ", err.Error())
		}

		if len(signingKeyBytes) > 0 {
			config.TokenSigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(signingKeyBytes)
			if err != nil {
				log.Warn("error parsing RSA private key from signing key bytes: ", err.Error())
			}
		}
	}

	if config.TokenSigningKey == nil {
		log.Warnf("signing key not provided: using default signing key")
		config.TokenSigningKey = oauth.DefaultSigningKey
	}
}

func initTokenVerifyingKey(config *server.Config, verifyingKeyPath *string) {
	if verifyingKeyPath != nil && *verifyingKeyPath != "" {
		verifyingKeyBytes, err := ioutil.ReadFile(*verifyingKeyPath)
		if err != nil {
			log.Warn("error reading verifying key file: ", err.Error())
		}

		if len(verifyingKeyBytes) > 0 {
			config.TokenVerifyingKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyingKeyBytes)
			if err != nil {
				log.Warn("error parsing RSA public key from verifying key bytes: ", err.Error())
			}
		}
	}

	if config.TokenVerifyingKey == nil {
		log.Warnf("verifying key not provided: using default verifying key")
		config.TokenVerifyingKey = oauth.DefaultVerifyingKey
	}
}

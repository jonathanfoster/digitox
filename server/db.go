package server

import (
	"os"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // nolint: golint
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/store"
)

// InitDB initializes the database.
func InitDB(dataSource string) error {
	if _, err := os.Stat(dataSource); os.IsNotExist(err) {
		log.Warnf("data source %s does not exist: initializing empty data source", dataSource)
		dirname := path.Dir(dataSource)
		if err := os.MkdirAll(dirname, 0700); err != nil {
			return errors.Wrap(err, "error initializing data source directory")
		}

		f, err := os.OpenFile(dataSource, os.O_RDONLY|os.O_CREATE, 0600)
		defer f.Close() // nolint: errcheck, megacheck
		if err != nil {
			return errors.Wrap(err, "error initializing data source file")
		}
	}

	var err error
	store.DB, err = gorm.Open("sqlite3", dataSource)
	if err != nil {
		return errors.Wrap(err, "error initializing database connection")
	}

	store.DB.SetLogger(&store.GormLogger{})

	store.DB.AutoMigrate(&blocklist.Blocklist{})
	store.DB.AutoMigrate(&session.Session{})

	return nil
}

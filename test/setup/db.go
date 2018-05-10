package setup

import (
	"os"

	"github.com/jonathanfoster/digitox/server"
	log "github.com/sirupsen/logrus"
)

// TestDB initializes a test database in bin/test/sessions.db.
func TestDB() {
	if err := server.InitDB(os.Getenv("GOPATH")+"/src/github.com/jonathanfoster/digitox/bin/test/sessions.db", false); err != nil {
		log.Panic("error initializing test database: ", err.Error())
	}
}

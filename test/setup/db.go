package setup

import (
	"os"

	"github.com/jonathanfoster/digitox/server"
)

// TestDB initializes a test database in bin/test/sessions.db.
func TestDB() {
	server.InitDB(os.Getenv("GOPATH")+"/src/github.com/jonathanfoster/digitox/bin/test/sessions.db", false)
}

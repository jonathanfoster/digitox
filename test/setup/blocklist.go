package setup

import (
	"os"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/store"

	log "github.com/sirupsen/logrus"
)

// NewTestBlocklist creates a test blocklist instance.
func NewTestBlocklist() *blocklist.Blocklist {
	list := blocklist.New()
	list.Name = "test"
	list.Domains = []string{"www.reddit.com", "news.ycombinator.com"}

	return list
}

// TestBlocklistStore creates the test blocklist directory and initializes the blocklist store.
func TestBlocklistStore() {
	var dirname = os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/blocklists/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test blocklist directory %s: %s", dirname, err.Error())
	}

	store.Blocklist = store.NewFileStore(dirname)
}

// TestBlocklist creates and saves a test blocklist.
func TestBlocklist() *blocklist.Blocklist {
	list := NewTestBlocklist()

	if err := list.Save(); err != nil {
		log.Panic("error saving test blocklist: ", err.Error())
	}

	return list
}

package setup

import (
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
)

// NewTestBlocklist creates a test blocklist instance.
func NewTestBlocklist() *blocklist.Blocklist {
	list := blocklist.New()
	list.Name = "test"
	list.Domains = []string{"www.reddit.com", "news.ycombinator.com"}

	return list
}

// TestBlocklist creates and saves a test blocklist.
func TestBlocklist() *blocklist.Blocklist {
	list := NewTestBlocklist()

	if err := list.Save(); err != nil {
		log.Panic("error saving test blocklist: ", err.Error())
	}

	return list
}

package setup

import (
	"os"
	"time"

	"github.com/jonathanfoster/digitox/models/blocklist"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/store"
)

// NewTestSession creates a test session instance with a specific blocklist ID and device name.
func NewTestSession() *session.Session {
	now := time.Now().UTC()
	sess := session.New()
	sess.Name = "test"
	sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	sess.RepeatEveryDay()
	sess.Blocklists = []*blocklist.Blocklist{NewTestBlocklist()}

	return sess
}

// TestSessionStore creates the test session directory and initializes the session store.
func TestSessionStore() {
	var dirname = os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test session directory %s: %s", dirname, err.Error())
	}

	store.Session = store.NewFileStore(dirname)
	TestDB()
}

// TestSession creates and saves a test session with a specific blocklist ID.
func TestSession() *session.Session {
	sess := NewTestSession()

	if err := sess.Save(); err != nil {
		log.Panicf("error saving test session %s: %s", sess.ID.String(), err.Error())
	}

	return sess
}

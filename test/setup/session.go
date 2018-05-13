package setup

import (
	"time"

	"github.com/jonathanfoster/digitox/models/blocklist"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/session"
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

// TestSession creates and saves a test session with a specific blocklist ID.
func TestSession() *session.Session {
	sess := NewTestSession()
	list := sess.Blocklists[0]

	if err := list.Save(); err != nil {
		log.Panicf("error saving test blocklist %s: %s", list.ID.String(), err.Error())
	}

	if err := sess.Save(); err != nil {
		log.Panicf("error saving test session %s: %s", sess.ID.String(), err.Error())
	}

	return sess
}

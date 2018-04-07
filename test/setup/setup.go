package setup

import (
	"os"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/store"
)

// TestBlocklistDirname creates and sets the test blocklist directory.
func TestBlocklistDirname() {
	dirname := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/freedom/bin/test/blocklists/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test blocklist directory %s: %s", dirname, err.Error())
	}

	store.Blocklist = store.NewFileStore(dirname)
}

// TestSessionDirname creates and sets the test session directory.
func TestSessionDirname() {
	dirname := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test session directory %s: %s", dirname, err.Error())
	}

	store.Session = store.NewFileStore(dirname)
}

// NewTestBlocklist creates a test blocklist instance.
func NewTestBlocklist() *blocklist.Blocklist {
	list := blocklist.New()
	list.Name = "test"
	list.Domains = []string{"www.reddit.com", "news.ycombinator.com"}

	return list
}

// TestBlocklist creates a test blocklist.
func TestBlocklist() *blocklist.Blocklist {
	list := NewTestBlocklist()

	if err := list.Save(); err != nil {
		log.Panic("error saving test blocklist: ", err.Error())
	}

	return list
}

// NewTestSession creates a test session instance.
func NewTestSession() *session.Session {
	now := time.Now().UTC()
	sess := session.New()
	sess.Name = "test"
	sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 11, 59, 59, 0, now.Location())
	sess.RepeatEveryDay()
	sess.Blocklists = []uuid.UUID{
		uuid.NewV4(),
	}

	return sess
}

// TestSession creates and saves a test session.
func TestSession() *session.Session {
	return TestSessionWithBlocklist(uuid.UUID{})
}

// TestSessionWithBlocklist creates and saves a test session with a specific blocklist ID.
func TestSessionWithBlocklist(list uuid.UUID) *session.Session {
	sess := NewTestSession()
	empty := uuid.UUID{}

	if list != empty {
		sess.Blocklists = []uuid.UUID{list}
	}

	if err := sess.Save(); err != nil {
		log.Panicf("error saving test session %s: %s", sess.ID.String(), err.Error())
	}

	return sess
}

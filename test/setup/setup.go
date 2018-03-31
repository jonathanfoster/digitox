package setup

import (
	"os"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/store"
	"github.com/jonathanfoster/freedom/store/fs"
)

// TestBlocklistDirname creates and sets the test blocklist directory.
func TestBlocklistDirname() {
	dirname := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test blocklist directory %s: %s", dirname, err.Error())
	}

	store.Blocklist = fs.NewFileStore(dirname)
}

// TestSessionDirname creates and sets the test session directory.
func TestSessionDirname() {
	dirname := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test session directory %s: %s", dirname, err.Error())
	}

	store.Session = fs.NewFileStore(dirname)
}

// TestBlocklist creates a test blocklist.
func TestBlocklist() *blocklist.Blocklist {
	testlist := blocklist.New()
	testlist.Name = "test"
	testlist.Hosts = append(testlist.Hosts, "www.reddit.com")

	if err := testlist.Save(); err != nil {
		log.Panic("error saving test blocklist: ", err.Error())
	}

	return testlist
}

// NewTestSession creates a test session instance.
func NewTestSession() *session.Session {
	sess := session.New()
	sess.Name = "test"
	sess.Starts = time.Now()
	sess.Ends = sess.Starts.Add(time.Hour * 1)
	sess.Repeats = []session.RepeatSchedule{
		session.EverySunday,
		session.EveryMonday,
		session.EveryTuesday,
		session.EveryWednesday,
		session.EveryThursday,
		session.EveryFriday,
		session.EverySaturday,
	}
	sess.Blocklists = []blocklist.Blocklist{
		*blocklist.New(),
	}

	return sess
}

// TestSession creates and saves a test session.
func TestSession() *session.Session {
	sess := NewTestSession()
	sess.ID = uuid.NewV4()

	if err := sess.Save(); err != nil {
		log.Panicf("error saving test session %s: %s", sess.ID.String(), err.Error())
	}

	return sess
}

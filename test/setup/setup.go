package setup

import (
	"os"
	"time"

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
		log.Fatalf("error creating test blocklist directory %s: %s", dirname, err.Error())
	}

	store.Blocklist = fs.NewFileStore(dirname)
}

// TestSessionDirname creates and sets the test session directory.
func TestSessionDirname() {
	dirname := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Fatalf("error creating test session directory %s: %s", dirname, err.Error())
	}

	store.Session = fs.NewFileStore(dirname)
}

// TestBlocklist creates a test blocklist.
func TestBlocklist() *blocklist.Blocklist {
	testlist := blocklist.New()
	testlist.Name = "test"
	testlist.Hosts = append(testlist.Hosts, "www.reddit.com")

	if err := testlist.Save(); err != nil {
		log.Fatal("error saving test blocklist: ", err.Error())
	}

	return testlist
}

// NewTestSession creates a test session instance.
func NewTestSession() *session.Session {
	testsess := session.New()
	testsess.Name = "test"
	testsess.Starts = time.Now()
	testsess.Ends = testsess.Starts.Add(time.Hour * 1)
	testsess.Repeats = []session.RepeatSchedule{
		session.EverySunday,
		session.EveryMonday,
		session.EveryTuesday,
		session.EveryWednesday,
		session.EveryThursday,
		session.EveryFriday,
		session.EverySaturday,
	}
	testsess.Blocklists = []blocklist.Blocklist{
		*blocklist.New(),
	}

	return testsess
}

// TestSession creates and saves a test session.
func TestSession() *session.Session {
	testsess := NewTestSession()

	if err := testsess.Save(); err != nil {
		log.Fatal("error saving test session: ", err.Error())
	}

	return testsess
}

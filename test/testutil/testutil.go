package testutil

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/store"
)

// SetTestBlocklistDirname creates and sets the test blocklist directory.
func SetTestBlocklistDirname() error {
	gopath := os.Getenv("GOPATH")
	dirname := gopath + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		return errors.Wrapf(err, "error creating test blocklist directory %s", dirname)
	}

	store.Blocklist.SetDirname(dirname)

	return nil
}

// SetTestSessionDirname creates and sets the test session directory.
func SetTestSessionDirname() error {
	gopath := os.Getenv("GOPATH")
	dirname := gopath + "/src/github.com/jonathanfoster/freedom/bin/test/sessions/"

	if err := os.MkdirAll(dirname, 0700); err != nil {
		return errors.Wrapf(err, "error creating test session directory %s", dirname)
	}

	store.Session.SetDirname(dirname)

	return nil
}

// SaveTestBlocklist creates a test blocklist.
func SaveTestBlocklist() (*blocklist.Blocklist, error) {
	testlist := blocklist.New(uuid.NewV4().String())
	testlist.Name = "test"
	testlist.Hosts = append(testlist.Hosts, "www.reddit.com")

	if err := testlist.Save(); err != nil {
		return nil, err
	}

	return testlist, nil
}

// NewTestSession creates a test session instance
func NewTestSession() *session.Session {
	testsess := session.New(uuid.NewV4().String())
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
	testsess.Blocklists = []string{uuid.NewV4().String()}
	testsess.Devices = []string{uuid.NewV4().String()}

	return testsess
}

// SaveTestSession creates and saves a test session.
func SaveTestSession() (*session.Session, error) {
	testsess := NewTestSession()

	if err := testsess.Save(); err != nil {
		return nil, err
	}

	return testsess, nil
}

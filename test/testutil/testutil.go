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

	blocklist.Dirname = dirname
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

	session.Dirname = dirname
	store.Session.SetDirname(dirname)

	return nil
}

// CreateTestBlocklist creates a test blocklist.
func CreateTestBlocklist() (*blocklist.Blocklist, error) {
	testlist := blocklist.New(uuid.NewV4().String())
	testlist.Name = "test"
	testlist.Hosts = append(testlist.Hosts, "www.reddit.com")

	if err := testlist.Save(); err != nil {
		return nil, err
	}

	return testlist, nil
}

// CreateTestSession creates a test session.
func CreateTestSession() (*session.Session, error) {
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

	if err := testsess.Save(); err != nil {
		return nil, err
	}

	return testsess, nil
}

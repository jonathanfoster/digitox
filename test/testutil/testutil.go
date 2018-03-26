package testutil

import (
	"os"
	"path/filepath"

	"github.com/jonathanfoster/freedom/model/blocklist"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// SetTestBlocklistDirname creates and sets the test blocklist directory.
func SetTestBlocklistDirname() error {
	path := "../../bin/test/blocklists/"
	dirname, err := filepath.Abs(path)
	if err != nil {
		return errors.Wrapf(err, "error returning absolute test blocklist directory %s", path)
	}

	if err := os.MkdirAll(path, 0644); err != nil {
		return errors.Wrapf(err, "error creating test blocklist directory %s", path)
	}

	blocklist.Dirname = dirname

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

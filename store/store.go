package store

import (
	"github.com/pkg/errors"
)

var (
	// Blocklist is the default blocklist store.
	Blocklist Interface
	// Session is the default session store.
	Session Interface

	// ErrNotExist is the error returned when a record does not exist.
	ErrNotExist = errors.New("record does not exist")
)

func init() {
	Blocklist = NewFile("/etc/freedom/blocklists/")
	Session = NewFile("/etc/freedom/sessions/")
}

// Interface represents the store interface.
type Interface interface {
	All() ([]string, error)
	Find(id string, out interface{}) error
	Remove(id string) error
	Save(id string, v interface{}) error
	SetDirname(dirname string)
}

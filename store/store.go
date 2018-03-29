package store

import (
	"github.com/pkg/errors"
)

var (
	// Blocklist is the blocklist storage implementation.
	Blocklist Interface
	// Session is the session storage implementation.
	Session Interface

	// ErrNotExist is the error returned when a record does not exist.
	ErrNotExist = errors.New("record does not exist")
)

// Interface is the session storage interface.
type Interface interface {
	All() ([]string, error)
	Find(id string, out interface{}) error
	Remove(id string) error
	Save(id string, value interface{}) error
}

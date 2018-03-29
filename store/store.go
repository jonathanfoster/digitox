package store

import (
	"github.com/pkg/errors"
)

var (
	// Blocklist is the blocklist storage implementation.
	Blocklist BlocklistStore
	// Session is the session storage implementation.
	Session SessionStore

	// ErrNotExist is the error returned when a record does not exist.
	ErrNotExist = errors.New("record does not exist")
)

// BlocklistStore is the blocklist storage interface.
type BlocklistStore interface {
	All() ([]string, error)
	Find(id string, out interface{}) error
	Remove(id string) error
	Save(id string, v interface{}) error
}

// SessionStore is the session storage interface.
type SessionStore interface {
	All() ([]string, error)
	Find(id string, out interface{}) error
	Remove(id string) error
	Save(id string, v interface{}) error
}

package store

import (
	"github.com/pkg/errors"
)

var (
	// Blocklist is the blocklist storage implementation.
	Blocklist Interface
	// Device is the device storage implementation.
	Device Interface
	// Session is the session storage implementation.
	Session Interface

	// ErrNotFound is the error returned when a record is not found.
	ErrNotFound = errors.New("record not found")
)

// Interface is the storage interface.
type Interface interface {
	All() ([]string, error)
	Exists(id string) (bool, error)
	Find(id string, out interface{}) error
	Init() error
	Remove(id string) error
	Save(id string, value interface{}) error
}

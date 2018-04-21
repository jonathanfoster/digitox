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

	// ErrNotExist is the error returned when a record does not exist.
	ErrNotExist = errors.New("record does not exist")
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

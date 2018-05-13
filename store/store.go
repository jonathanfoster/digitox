package store

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	// DB is the global database instance.
	DB *gorm.DB
	// Device is the device storage implementation.
	Device Interface

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

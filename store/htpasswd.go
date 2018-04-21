package store

import (
	"os"

	"github.com/foomo/htpasswd"
	"github.com/pkg/errors"
)

func init() {
	Device = NewHtpasswdStore("/etc/digitox/passwd")
}

// Credentials represents an interface for storing a record in htpasswd file.
type Credentials interface {
	Username() string
	Password() string
	Hash() string
}

// HtpasswdStore represents a htpasswd store.
type HtpasswdStore struct {
	Filename string
}

// NewHtpasswdStore creates a HtpasswdStore instance.
func NewHtpasswdStore(filename string) *HtpasswdStore {
	return &HtpasswdStore{
		Filename: filename,
	}
}

// All retrieves all devices from htpasswd file.
func (h *HtpasswdStore) All() ([]string, error) {
	credentials, err := htpasswd.ParseHtpasswdFile(h.Filename)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing htpasswd file")
	}

	deviceNames := make([]string, len(credentials))

	i := 0
	for k := range credentials {
		deviceNames[i] = k
		i++
	}

	return deviceNames, nil
}

// Exists checks if a device exists in htpasswd file.
func (h *HtpasswdStore) Exists(name string) (bool, error) {
	return false, nil
}

// Find finds a device by name in htpasswd file.
func (h *HtpasswdStore) Find(name string, out interface{}) error {
	passwords, err := htpasswd.ParseHtpasswdFile(h.Filename)
	if err != nil {
		return errors.Wrapf(err, "error finding device %s in htpasswd file", name)
	}

	out, ok := passwords[name]
	if !ok {
		return ErrNotExist
	}

	return nil
}

// Init creates the htpasswd file directory and all parent directories.
func (h *HtpasswdStore) Init() error {
	if err := os.MkdirAll(h.Filename, 0700); err != nil {
		return errors.Wrapf(err, "error initializing directory %s", h.Filename)
	}

	return nil
}

// Remove removes device from htpasswd file.
func (h *HtpasswdStore) Remove(name string) error {
	return nil
}

// Save writes device to htpasswd file.
func (h *HtpasswdStore) Save(name string, v interface{}) error {
	credentials, ok := v.(Credentials)
	if !ok {
		return errors.New("value is not of type credentials")
	}

	if credentials.Hash() != "" {
		if err := htpasswd.SetPasswordHash(h.Filename, credentials.Username(), credentials.Hash()); err != nil {
			return errors.Wrapf(err, "error saving device %s hash", credentials.Username())
		}

	} else {
		if err := htpasswd.SetPassword(h.Filename, credentials.Username(), credentials.Password(), htpasswd.HashBCrypt); err != nil {
			return errors.Wrapf(err, "error saving device %s password", credentials.Username())
		}

	}

	return nil
}

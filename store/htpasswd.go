package store

import (
	"os"

	"github.com/foomo/htpasswd"
	"github.com/pkg/errors"
)

func init() {
	Device = NewHtpasswdStore("/etc/digitox/passwd")
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

// Exists checks if a file exists without reading the entire file.
func (h *HtpasswdStore) Exists(id string) (bool, error) {
	return false, nil
}

// Find finds a device by ID in htpasswd.
func (h *HtpasswdStore) Find(id string, out interface{}) error {
	return nil
}

// Init creates the htpasswd store directory and all parent directories.
func (h *HtpasswdStore) Init() error {
	if err := os.MkdirAll(h.Filename, 0700); err != nil {
		return errors.Wrapf(err, "error initializing directory %s", h.Filename)
	}

	return nil
}

// Remove removes the device from htpasswd.
func (h *HtpasswdStore) Remove(id string) error {
	return nil
}

// Save writes value to htpasswd.
func (h *HtpasswdStore) Save(id string, v interface{}) error {
	return nil
}

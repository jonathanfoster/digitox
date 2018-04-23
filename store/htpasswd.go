package store

import (
	"os"
	"path"
	"reflect"
	"runtime"

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

// Exists checks if a device exists in htpasswd file.
func (h *HtpasswdStore) Exists(name string) (bool, error) {
	passwords, err := htpasswd.ParseHtpasswdFile(h.Filename)
	if err != nil {
		return false, errors.Wrapf(err, "error checking if device %s exists in htpasswd file", name)
	}

	_, ok := passwords[name]

	return ok, nil
}

// Find finds a device by name in htpasswd file.
func (h *HtpasswdStore) Find(name string, out interface{}) (err error) {
	passwords, err := htpasswd.ParseHtpasswdFile(h.Filename)
	if err != nil {
		return errors.Wrapf(err, "error finding device %s in htpasswd file", name)
	}

	hash, ok := passwords[name]
	if !ok {
		return ErrNotExist
	}

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}

			err = r.(error)
		}
	}()

	// Assumes out is a pointer to a struct with Name and Hash fields
	// If not the above defer func will catch and set err value
	ptr := reflect.ValueOf(out)
	value := ptr.Elem()

	value.FieldByName("Name").SetString(name)
	value.FieldByName("Hash").SetString(hash)

	return nil
}

// Init creates the htpasswd file directory and all parent directories.
func (h *HtpasswdStore) Init() error {
	dirname := path.Dir(h.Filename)

	if err := os.MkdirAll(dirname, 0700); err != nil {
		return errors.Wrapf(err, "error initializing htpasswd directory %s", dirname)
	}

	f, err := os.OpenFile(h.Filename, os.O_RDONLY|os.O_CREATE, 0600)
	defer f.Close() // nolint: errcheck, megacheck

	if err != nil {
		return errors.Wrapf(err, "error initializing htpasswd file %s", h.Filename)
	}

	return nil
}

// Remove removes device from htpasswd file.
func (h *HtpasswdStore) Remove(name string) error {
	if err := htpasswd.RemoveUser(h.Filename, name); err != nil {
		if err == htpasswd.ErrNotExist {
			return ErrNotExist
		}

		return errors.Wrapf(err, "error removing device %s from htpasswd file", name)
	}

	return nil
}

// Save writes device to htpasswd file.
func (h *HtpasswdStore) Save(name string, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}

			err = r.(error)
		}
	}()

	// Assumes v is a pointer to a struct with Name, Password, and Hash fields
	// If not the above defer func will catch and set err value
	ptr := reflect.ValueOf(v)
	value := ptr.Elem()

	password := value.FieldByName("Password").String()
	hash := value.FieldByName("Hash").String()

	if password != "" {
		if err := htpasswd.SetPassword(h.Filename, name, password, htpasswd.HashBCrypt); err != nil {
			return errors.Wrapf(err, "error saving device %s: error setting password", name)
		}
	} else {
		if err := htpasswd.SetPasswordHash(h.Filename, name, hash); err != nil {
			return errors.Wrapf(err, "error saving device %s: error setting password hash", name)
		}
	}

	return nil
}

package store

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// File represents a file store.
type File struct {
	Dirname string
}

// NewFile creates a File instance.
func NewFile(dirname string) *File {
	return &File{
		Dirname: dirname,
	}
}

// All retrieves all values from file system.
func (f *File) All() ([]string, error) {
	files, err := ioutil.ReadDir(f.Dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrap(err, "error retrieving all values")
	}

	var filenames []string

	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames, nil
}

// Find finds a blocklist by ID in the filesystem.
func (f *File) Find(id string, out interface{}) error {
	filename, err := JoinPath(id, f.Dirname)
	if err != nil {
		return errors.Wrapf(err, "error creating file name to find %s", id)
	}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}

		return errors.Wrapf(err, "error reading file %s", filename)
	}

	if err := json.Unmarshal(buf, out); err != nil {
		return errors.Wrap(err, "error unmarshaling value")
	}

	return nil
}

// Remove removes the session from the filesystem.
func (f *File) Remove(id string) error {
	filename, err := JoinPath(id, f.Dirname)
	if err != nil {
		return errors.Wrapf(err, "error creating file name to remove %s", id)
	}

	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}

		return errors.Wrapf(err, "error removing file %s", filename)
	}

	return nil
}

// Save writes value to the file system.
func (f *File) Save(id string, v interface{}) error {
	buf, err := json.Marshal(v)
	if err != nil {
		return errors.Wrapf(err, "error marshaling value %s", id)
	}

	filename, err := JoinPath(id, f.Dirname)
	if err != nil {
		return errors.Wrapf(err, "error creating file name for value %s", id)
	}

	if err := ioutil.WriteFile(filename, buf, 0700); err != nil {
		return errors.Wrapf(err, "error writing value file %s", filename)
	}

	return nil
}

// SetDirname sets the file system directory name.
func (f *File) SetDirname(dirname string) {
	f.Dirname = dirname
}

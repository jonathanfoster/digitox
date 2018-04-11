package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

func init() {
	Blocklist = NewFileStore("/etc/digitox/blocklists/")
	Session = NewFileStore("/etc/digitox/sessions/")
}

// FileStore represents a file store.
type FileStore struct {
	Dirname string
}

// NewFileStore creates a FileStore instance.
func NewFileStore(dirname string) *FileStore {
	return &FileStore{
		Dirname: dirname,
	}
}

// All retrieves all values from file system.
func (f *FileStore) All() ([]string, error) {
	files, err := ioutil.ReadDir(f.Dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrap(err, "error retrieving all values")
	}

	var vv []string

	for _, file := range files {
		if !file.IsDir() {
			vv = append(vv, file.Name())
		}
	}

	return vv, nil
}

// Exists checks if a file exists without reading the entire file.
func (f *FileStore) Exists(id string) (bool, error) {
	filename, err := JoinPath(id, f.Dirname)
	if err != nil {
		return false, errors.Wrapf(err, "error creating file name to check if %s exists", id)
	}

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.Wrapf(err, "error retrieving file info for file %s", filename)
	}

	return true, nil
}

// Find finds a blocklist by ID in the filesystem.
func (f *FileStore) Find(id string, out interface{}) error {
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

// Init creates the file store directory and all parent directories.
func (f *FileStore) Init() error {
	if err := os.MkdirAll(f.Dirname, 0700); err != nil {
		return errors.Wrapf(err, "error initializing directory %s", f.Dirname)
	}

	return nil
}

// Remove removes the session from the filesystem.
func (f *FileStore) Remove(id string) error {
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
func (f *FileStore) Save(id string, v interface{}) error {
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
func (f *FileStore) SetDirname(dirname string) {
	f.Dirname = dirname
}

// JoinPath sanitizes ID and joins with directory path to create a file path.
// The file path is then checked to ensure its directory is the directory path
// to prevent directory traversal using relative paths.
func JoinPath(filename string, dirname string) (string, error) {
	filename = path.Join(dirname, validator.SafeFileName(filename))
	filename, err := filepath.Abs(filename)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute path for %s", filename)
	}

	dirname, err = filepath.Abs(dirname)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute path for %s", dirname)
	}

	if filepath.Dir(filename) != dirname {
		return "", fmt.Errorf("file name %s not in expected directory %s", filename, dirname)
	}

	return filename, nil
}

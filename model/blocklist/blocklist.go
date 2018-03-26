package blocklist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/freedom/model"
)

var (
	// Dirname is the name of the blocklists directory.
	Dirname = "/etc/freedom/blocklists/"
	// ErrNotExist is the error returned when a blocklist does not exist.
	ErrNotExist = errors.New("blocklist does not exist")
)

// Blocklist represents a blocklist.
type Blocklist struct {
	ID    string   `json:"id" valid:"required, uuidv4"`
	Name  string   `json:"name"`
	Hosts []string `json:"hosts"`
}

// New creates a Blocklist instance.
func New(id string) *Blocklist {
	return &Blocklist{
		ID:    id,
		Hosts: []string{},
	}
}

// All retrieves all blocklists from filesystem.
func All() ([]*Blocklist, error) {
	files, err := ioutil.ReadDir(Dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrap(err, "error retrieving all blocklists")
	}

	lists := []*Blocklist{}

	for _, file := range files {
		if !file.IsDir() {
			lists = append(lists, New(file.Name()))
		}
	}

	return lists, nil
}

// FileName sanitizes ID and joins with blocklist directory path to create a
// file path. The file path is then checked to ensure its directory is the
// blocklist directory to prevent directory traversal using relative paths.
func FileName(id string) (string, error) {
	filename := path.Join(Dirname, validator.SafeFileName(id))
	filename, err := filepath.Abs(filename)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute blocklist file path %s", filename)
	}

	dirname, err := filepath.Abs(Dirname)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute blocklist directory path %s", Dirname)
	}

	if filepath.Dir(filename) != dirname {
		return "", fmt.Errorf("blocklist file path %s not in blocklist directory %s", filename, dirname)
	}

	return filename, nil
}

// Find finds a blocklist by ID in the filesystem.
func Find(id string) (*Blocklist, error) {
	filename, err := FileName(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating blocklist file name to find ID  %s", id)
	}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrapf(err, "error reading blocklist file %s", filename)
	}

	list := &Blocklist{}
	if err := json.Unmarshal(buf, list); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling blocklist")
	}

	return list, nil
}

// Remove removes the blocklist from the filesystem.
func Remove(id string) error {
	filename, err := FileName(id)
	if err != nil {
		return errors.Wrapf(err, "error creating blocklist file name to remove ID  %s", id)
	}

	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}

		return errors.Wrapf(err, "error removing blocklist file %s", filename)
	}

	return nil
}

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() error {
	if _, err := b.Validate(); err != nil {
		return model.NewValidatorError(fmt.Sprintf("error validating blocklist before save: %s", err.Error()))
	}

	buf, err := json.Marshal(b)
	if err != nil {
		return errors.Wrapf(err, "error marshaling blocklist %s", b.ID)
	}

	filename, err := FileName(b.ID)
	if err != nil {
		return errors.Wrapf(err, "error creating blocklist file name to save ID  %s", b.ID)
	}

	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		return errors.Wrapf(err, "error writing blocklist file %s", filename)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (b *Blocklist) Validate() (bool, error) {
	result, err := validator.ValidateStruct(b)
	if err != nil {
		err = model.NewValidatorError(err.Error())
	}

	return result, err
}

package blocklist

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	validator "github.com/asaskevich/govalidator"
	"github.com/jonathanfoster/freedom/model"
	"github.com/pkg/errors"
)

// Dirname is the name of the blocklists directory.
var Dirname = "/etc/squid/blocklists/"

// Blocklist represents a blocklist.
type Blocklist struct {
	Name    string   `json:"name" valid:"required"`
	Dirname string   `json:"dirname" valid:"required"`
	Hosts   []string `json:"hosts"`

	origName string
}

// New creates a Blocklist instance.
func New(name string) *Blocklist {
	return &Blocklist{
		Name:     name,
		Dirname:  Dirname,
		Hosts:    []string{},
		origName: name,
	}
}

// All retrieves all blocklists from filesystem.
func All() ([]*Blocklist, error) {
	files, err := ioutil.ReadDir(Dirname)
	if err != nil {
		return nil, err
	}

	lists := []*Blocklist{}

	for _, file := range files {
		if !file.IsDir() {
			lists = append(lists, New(file.Name()))
		}
	}

	return lists, nil
}

// Find finds a blocklist by name in the filesystem.
func Find(name string) (*Blocklist, error) {
	buf, err := ioutil.ReadFile(path.Join(Dirname, name))
	if err != nil {
		return nil, err
	}

	list := New(name)
	list.Hosts = strings.Split(string(buf), "\n")

	return list, nil
}

// Remove removes the blocklist from the filesystem.
func Remove(name string) error {
	return os.Remove(path.Join(Dirname, name))
}

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() error {
	if _, err := b.Validate(); err != nil {
		errors.Wrap(err, "")
		return err
	}

	var buffer bytes.Buffer

	for _, host := range b.Hosts {
		if _, err := buffer.WriteString(host + "\n"); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(path.Join(b.Dirname, b.Name), buffer.Bytes(), 0644); err != nil {
		return err
	}

	// Handle abandoned list when name changes
	if b.origName != "" && b.Name != b.origName {
		if err := os.Remove(path.Join(Dirname, b.origName)); err != nil {
			return err
		}
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

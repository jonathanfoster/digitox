package blocklist

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/freedom/model"
)

var (
	// Dirname is the name of the blocklists directory.
	Dirname = "/etc/squid/blocklists/"
	// ErrNotExist is the error returned when a blocklist does not exist.
	ErrNotExist = errors.New("blocklist does not exist")

	nameRegexp = regexp.MustCompile(`^#\s*name\s*:\s*(.*)\s*$`)
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

// Find finds a blocklist by ID in the filesystem.
func Find(id string) (*Blocklist, error) {
	filename := path.Join(Dirname, id)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrapf(err, "error reading blocklist file %s", filename)
	}

	list := &Blocklist{}
	if err := Unmarshal(buf, list); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling blocklist")
	}

	return list, nil
}

// Remove removes the blocklist from the filesystem.
func Remove(id string) error {
	filename := path.Join(Dirname, id)
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

	buf, err := Marshal(b)
	if err != nil {
		return errors.Wrapf(err, "error marshaling blocklist %s", b.ID)
	}

	filename := path.Join(Dirname, b.ID)
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

// Marshal returns encoding of v.
func Marshal(v *Blocklist) ([]byte, error) {
	var buffer bytes.Buffer

	if v.Name != "" {
		buffer.WriteString(fmt.Sprintf("# name: %s\n", v.Name))
	}

	for _, host := range v.Hosts {
		if _, err := buffer.WriteString(host + "\n"); err != nil {
			return nil, errors.Wrapf(err, "error marshaling blocklist: error writing host %s to buffer", host)
		}
	}

	return buffer.Bytes(), nil
}

// Unmarshal parses the encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v *Blocklist) error {
	s := string(data)
	lines := strings.Split(s, "\n")
	if len(lines) == 1 && lines[0] == "" {
		return errors.New("error unmarshaling blocklist: data is empty")
	}

	// nil is no match
	// match[0] is full match
	// match[1] is group 1 match which contains name value
	match := nameRegexp.FindStringSubmatch(lines[0])
	if len(match) == 2 {
		// Name match found, first line is name and all other lines are hosts
		v.Name = match[1]
		v.Hosts = lines[1:]
	} else {
		// No name match found, all lines are hosts
		v.Name = ""
		v.Hosts = lines
	}

	return nil
}

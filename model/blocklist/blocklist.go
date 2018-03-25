package blocklist

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Dirname is the name of the blocklists directory.
var Dirname = "/etc/squid/blocklists/"

// Blocklist represents a blocklist.
type Blocklist struct {
	Name    string   `json:"name"`
	Dirname string   `json:"dirname"`
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

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() error {
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
	if b.Name != b.origName {
		if err := os.Remove(path.Join(Dirname, b.origName)); err != nil {
			return err
		}
	}

	return nil
}

// Remove removes the blocklist from the filesystem.
func Remove(name string) error {
	return os.Remove(path.Join(Dirname, name))
}

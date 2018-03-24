package blocklist

import (
	"io/ioutil"
	"os"
	"path"
)

// Dirname is the name of the blocklists directory.
var Dirname = "/etc/squid/blocklists/"

// Blocklist represents a blocklist.
type Blocklist struct {
	Name    string `json:"name"`
	Dirname string `json:"dirname"`
}

// New creates a Blocklist instance.
func New() *Blocklist {
	return &Blocklist{}
}

// NewFromFile creates a Blocklist instance from file.
func NewFromFile(file *os.File) *Blocklist {
	return &Blocklist{
		Name:    file.Name(),
		Dirname: Dirname,
	}
}

// NewFromFileInfo creates a Blocklist instance from file info.
func NewFromFileInfo(file os.FileInfo) *Blocklist {
	return &Blocklist{
		Name:    file.Name(),
		Dirname: Dirname,
	}
}

// All retrieves all blocklists from filesystem.
func All() ([]*Blocklist, error) {
	files, err := ioutil.ReadDir(Dirname)
	if err != nil {
		return nil, err
	}

	list := []*Blocklist{}

	for _, file := range files {
		if !file.IsDir() {
			list = append(list, NewFromFileInfo(file))
		}
	}

	return list, nil
}

// Find finds a blocklist by name in the filesystem.
func Find(name string) (*Blocklist, error) {
	file, err := os.Open(path.Join(Dirname, name))
	if err != nil {
		return nil, err
	}

	return NewFromFile(file), nil
}

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() {
	// Write /etc/squid/blocklists/{id}
}

// Remove deletes the blocklist by name from the filesystem.
func Remove(name string) *Blocklist {
	// Delete /etc/squid/blocklists/{id}
	return nil
}

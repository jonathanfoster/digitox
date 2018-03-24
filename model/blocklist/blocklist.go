package blocklist

import (
	"io/ioutil"
)

// Dirname is the name of the blocklists directory.
var Dirname = "/etc/squid/blocklists/"

// Blocklist represents a blocklist.
type Blocklist struct {
	Name    string `json:"name"`
	Dirname string `json:"dirname"`
}

// New creates a Blocklist instance
func New() *Blocklist {
	return &Blocklist{}
}

// All retrieves all blocklists from filesystem.
func All() ([]Blocklist, error) {
	files, err := ioutil.ReadDir(Dirname)
	if err != nil {
		return nil, err
	}

	list := []Blocklist{}

	for _, file := range files {
		if !file.IsDir() {
			list = append(list, Blocklist{
				Name:    file.Name(),
				Dirname: Dirname,
			})
		}
	}

	return list, nil
}

// Find finds a blocklist by name in the filesystem.
func Find(name string) *Blocklist {
	// Search /etc/squid/blocklists/
	return nil
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

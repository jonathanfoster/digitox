package blocklist

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jonathanfoster/digitox/store"
)

// Blocklist represents a list of websites to block.
type Blocklist struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Domains []string  `json:"domains" valid:"required"`
}

// New creates a Blocklist instance.
func New() *Blocklist {
	return &Blocklist{
		ID:      uuid.NewV4(),
		Domains: []string{},
	}
}

// All retrieves all blocklists.
func All() ([]*Blocklist, error) {
	ff, err := store.Blocklist.All()
	if err != nil {
		return nil, err
	}

	var bb []*Blocklist

	for _, f := range ff {
		b, err := Find(f)
		if err != nil {
			return nil, err
		}

		bb = append(bb, b)
	}

	return bb, nil
}

// Exists checks if a blocklist exists by ID.
func Exists(id string) (bool, error) {
	exists, err := store.Blocklist.Exists(id)
	if err != nil {
		if err == store.ErrNotFound {
			return false, nil
		}

		return false, errors.Wrapf(err, "error checking if blocklist %s exists", id)
	}

	return exists, nil
}

// Find finds a blocklist by ID.
func Find(id string) (*Blocklist, error) {
	var list Blocklist

	if err := store.Blocklist.Find(id, &list); err != nil {
		return nil, errors.Wrapf(err, "error finding blocklist %s", id)
	}

	return &list, nil
}

// Remove removes the blocklist.
func Remove(id string) error {
	if err := store.Blocklist.Remove(id); err != nil {
		return errors.Wrapf(err, "error removing blocklist %s", id)
	}

	return nil
}

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() error {
	if err := store.Blocklist.Save(b.ID.String(), b); err != nil {
		return errors.Wrapf(err, "error saving blocklist %s", b.ID)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (b *Blocklist) Validate() (bool, error) {
	return validator.ValidateStruct(b)
}

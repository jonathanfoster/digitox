package blocklist

import (
	"fmt"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/freedom/models"
	"github.com/jonathanfoster/freedom/store"
)

// Blocklist represents a list of websites to block.
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
	if _, err := b.Validate(); err != nil {
		return models.NewValidatorError(fmt.Sprintf("error validating blocklist before save: %s", err.Error()))
	}

	if err := store.Blocklist.Save(b.ID, b); err != nil {
		return errors.Wrapf(err, "error saving blocklist %s", b.ID)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (b *Blocklist) Validate() (bool, error) {
	result, err := validator.ValidateStruct(b)
	if err != nil {
		err = models.NewValidatorError(err.Error())
	}

	return result, err
}

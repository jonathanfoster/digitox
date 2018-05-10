package blocklist

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jonathanfoster/digitox/store"
)

// Blocklist represents a list of websites to block.
type Blocklist struct {
	ID        uuid.UUID  `json:"id" gorm:"type:text"`
	Name      string     `json:"name"`
	Domains   domainList `json:"domains" valid:"required" gorm:"type:text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type domainList []string

func (d domainList) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	for _, domain := range d {
		buffer.WriteString(fmt.Sprintf("%s,", domain))
	}

	return buffer.String(), nil
}

func (d *domainList) Scan(value interface{}) error {
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return errors.Wrap(err, "error converting scanner value to string type")
	}

	b, ok := sv.([]byte)
	if !ok {
		return errors.New("scanner value is not byte array type")
	}

	s := string(b)

	for _, domain := range strings.Split(s, ",") {
		if domain != "" {
			*d = append(*d, domain)
		}
	}

	return nil
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
	var lists []*Blocklist

	if err := store.DB.Find(&lists).Error; err != nil {
		return nil, errors.Wrap(err, "error retrieving all blocklists")
	}

	return lists, nil
}

// Exists checks if a blocklist exists by ID.
func Exists(id uuid.UUID) (bool, error) {
	if _, err := Find(id); err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return false, nil
		}

		return false, errors.Wrap(err, "error checking if blocklist exists")
	}

	return true, nil
}

// Find finds a blocklist by ID.
func Find(id uuid.UUID) (*Blocklist, error) {
	var list Blocklist

	if err := store.DB.Find(&list, &Blocklist{ID: id}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return nil, errors.Wrap(err, "error finding blocklist")
	}

	return &list, nil
}

// Remove removes the blocklist.
func Remove(id uuid.UUID) error {
	exists, err := Exists(id)
	if err != nil {
		return errors.Wrap(err, "error removing blocklist")
	}

	if !exists {
		return store.ErrNotFound
	}

	if err := store.DB.Delete(&Blocklist{ID: id}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return errors.Wrap(err, "error removing blocklist")
	}

	return nil
}

// Save writes the blocklist to the filesystem.
func (b *Blocklist) Save() error {
	if err := store.DB.Save(b).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return errors.Wrap(err, "error saving blocklist")
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (b *Blocklist) Validate() (bool, error) {
	return validator.ValidateStruct(b)
}

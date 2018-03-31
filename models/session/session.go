package session

import (
	"fmt"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jonathanfoster/freedom/models"
	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/store"
)

// RepeatSchedule represents a scheduled repetition of a session.
type RepeatSchedule int

// Repeat schedule for each day of the week.
const (
	EverySunday RepeatSchedule = iota
	EveryMonday
	EveryTuesday
	EveryWednesday
	EveryThursday
	EveryFriday
	EverySaturday
)

// Session represents a time frame in which websites are blocked
type Session struct {
	ID         uuid.UUID             `json:"id"`
	Name       string                `json:"name"`
	Starts     time.Time             `json:"starts" valid:"required"`
	Ends       time.Time             `json:"ends" valid:"required"`
	Repeats    []RepeatSchedule      `json:"repeats"`
	Blocklists []blocklist.Blocklist `json:"blocklists" valid:"required"`
}

// New creates a Session instance.
func New() *Session {
	return &Session{
		ID:         uuid.NewV4(),
		Repeats:    []RepeatSchedule{},
		Blocklists: []blocklist.Blocklist{},
	}
}

// All retrieves all sessions from file system.
func All() ([]*Session, error) {
	ff, err := store.Session.All()
	if err != nil {
		return nil, err
	}

	var ss []*Session

	for _, f := range ff {
		s, err := Find(f)
		if err != nil {
			return nil, err
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// Find finds a session by ID.
func Find(id string) (*Session, error) {
	var sess Session

	if err := store.Session.Find(id, &sess); err != nil {
		return nil, errors.Wrapf(err, "error finding session %s", id)
	}

	return &sess, nil
}

// Remove removes the session from the file system.
func Remove(id string) error {
	if err := store.Session.Remove(id); err != nil {
		return errors.Wrapf(err, "error removing session %s", id)
	}

	return nil
}

// Save writes the session to the file system.
func (s *Session) Save() error {
	if _, err := s.Validate(); err != nil {
		return models.NewValidatorError(fmt.Sprintf("error validating session before save: %s", err.Error()))
	}

	if err := store.Session.Save(s.ID.String(), s); err != nil {
		return errors.Wrapf(err, "error saving session %s", s.ID)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (s *Session) Validate() (bool, error) {
	result, err := validator.ValidateStruct(s)
	if err != nil {
		err = models.NewValidatorError(err.Error())
	}

	return result, err
}

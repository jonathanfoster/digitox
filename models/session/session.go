package session

import (
	"fmt"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/freedom/models"
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

var (
	// Dirname is the name of the sessions directory.
	Dirname = "/etc/freedom/sessions/"
)

// Session represents a timeframe in which websites are blocked
type Session struct {
	ID         string           `json:"id" valid:"required, uuidv4"`
	Name       string           `json:"name"`
	Starts     time.Time        `json:"starts"`
	Ends       time.Time        `json:"ends"`
	Repeats    []RepeatSchedule `json:"repeats"`
	Blocklists []string         `json:"blocklists"`
	Devices    []string         `json:"devices"`
}

// New creates a Session instance.
func New(id string) *Session {
	return &Session{
		ID:         id,
		Repeats:    []RepeatSchedule{},
		Blocklists: []string{},
		Devices:    []string{},
	}
}

// All retrieves all sessions from file system.
func All() ([]*Session, error) {
	filesnames, err := store.Session.All()
	if err != nil {
		return nil, err
	}

	var list []*Session

	for _, filename := range filesnames {
		list = append(list, New(filename))
	}

	return list, nil
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

	if err := store.Session.Save(s.ID, s); err != nil {
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

package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/freedom/model"
	"github.com/jonathanfoster/freedom/model/pathutil"
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
	// ErrNotExist is the error returned when a session does not exist.
	ErrNotExist = errors.New("session does not exist")
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

// All retrieves all blocklists from filesystem.
func All() ([]*Session, error) {
	files, err := ioutil.ReadDir(Dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotExist
		}

		return nil, errors.Wrap(err, "error retrieving all sessions")
	}

	sessions := []*Session{}

	for _, file := range files {
		if !file.IsDir() {
			sessions = append(sessions, New(file.Name()))
		}
	}

	return sessions, nil
}

// Remove removes the session from the filesystem.
func Remove(id string) error {
	filename, err := pathutil.FileName(id, Dirname)
	if err != nil {
		return errors.Wrapf(err, "error creating session file name to remove ID  %s", id)
	}

	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}

		return errors.Wrapf(err, "error removing session file %s", filename)
	}

	return nil
}

// Save writes the session to the filesystem.
func (s *Session) Save() error {
	if _, err := s.Validate(); err != nil {
		return model.NewValidatorError(fmt.Sprintf("error validating session before save: %s", err.Error()))
	}

	buf, err := json.Marshal(s)
	if err != nil {
		return errors.Wrapf(err, "error marshaling session %s", s.ID)
	}

	filename, err := pathutil.FileName(s.ID, Dirname)
	if err != nil {
		return errors.Wrapf(err, "error creating session file name to save ID  %s", s.ID)
	}

	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		return errors.Wrapf(err, "error writing session file %s", filename)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (s *Session) Validate() (bool, error) {
	result, err := validator.ValidateStruct(s)
	if err != nil {
		err = model.NewValidatorError(err.Error())
	}

	return result, err
}

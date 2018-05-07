package session

import (
	"fmt"
	"strings"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/store"
)

// Session represents a time frame in which websites are blocked
type Session struct {
	ID             uuid.UUID   `json:"id" gorm:"type:text"`
	Name           string      `json:"name"`
	Starts         time.Time   `json:"starts" valid:"required"`
	Ends           time.Time   `json:"ends" valid:"required"`
	Blocklists     []uuid.UUID `json:"blocklists" valid:"required" gorm:"-"`
	EverySunday    bool        `json:"every_sunday"`
	EveryMonday    bool        `json:"every_monday"`
	EveryTuesday   bool        `json:"every_tuesday"`
	EveryWednesday bool        `json:"every_wednesday"`
	EveryThursday  bool        `json:"every_thursday"`
	EveryFriday    bool        `json:"every_friday"`
	EverySaturday  bool        `json:"every_saturday"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      *time.Time  `json:"deleted_at"`
}

// New creates a Session instance.
func New() *Session {
	return &Session{
		ID: uuid.NewV4(),
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

// Exists checks if a session exists by ID.
func Exists(id string) (bool, error) {
	exists, err := store.Session.Exists(id)
	if err != nil {
		if err == store.ErrNotFound {
			return false, nil
		}

		return false, errors.Wrapf(err, "error checking if session %s exists", id)
	}

	return exists, nil
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

// Active determines whether a session is active based on starts, ends, and daily repeat options.
func (s *Session) Active() bool {
	now := time.Now().UTC()

	starts := s.Starts.UTC()
	ends := s.Ends.UTC()

	if s.RepeatsToday() {
		starts = time.Date(now.Year(), now.Month(), now.Day(), starts.Hour(), starts.Minute(), starts.Second(), starts.Nanosecond(), starts.Location())
		ends = time.Date(now.Year(), now.Month(), now.Day(), ends.Hour(), ends.Minute(), ends.Second(), ends.Nanosecond(), ends.Location())
	}

	var active bool

	if starts.Before(now) && ends.After(now) {
		active = true
	} else {
		active = false
	}

	fields := log.Fields{
		"id":     s.ID,
		"starts": s.Starts.String(),
		"ends":   s.Ends.String(),
		"now":    now.String(),
		"active": active,
	}

	if active {
		log.WithFields(fields).Debugf("session %s is active", s.ID)
	} else {
		log.WithFields(fields).Debugf("session %s is not active", s.ID)
	}

	return active
}

// RepeatEveryDay sets session to repeat every day of the week.
func (s *Session) RepeatEveryDay() {
	s.EverySunday = true
	s.EveryMonday = true
	s.EveryTuesday = true
	s.EveryWednesday = true
	s.EveryThursday = true
	s.EveryFriday = true
	s.EverySaturday = true
}

// RepeatNever sets session to never repeat.
func (s *Session) RepeatNever() {
	s.EverySunday = false
	s.EveryMonday = false
	s.EveryTuesday = false
	s.EveryWednesday = false
	s.EveryThursday = false
	s.EveryFriday = false
	s.EverySaturday = false
}

// RepeatsToday returns true if the session repeats on today's day of the week.
func (s *Session) RepeatsToday() bool { // nolint: gocyclo
	now := time.Now().UTC()
	weekday := now.Weekday()

	return (weekday == time.Sunday && s.EverySunday) ||
		(weekday == time.Monday && s.EveryMonday) ||
		(weekday == time.Tuesday && s.EveryTuesday) ||
		(weekday == time.Wednesday && s.EveryWednesday) ||
		(weekday == time.Thursday && s.EveryThursday) ||
		(weekday == time.Friday && s.EveryFriday) ||
		(weekday == time.Saturday && s.EverySaturday)
}

// Save writes the session to the file system.
func (s *Session) Save() error {
	if err := store.Session.Save(s.ID.String(), s); err != nil {
		return errors.Wrapf(err, "error saving session %s", s.ID)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (s *Session) Validate() (bool, error) {
	var msgs []string

	listsExist := true

	// Validate blocklists exist
	for _, id := range s.Blocklists {
		exists, err := blocklist.Exists(id.String())
		if err != nil {
			return false, errors.Wrapf(err, "error validating session blocklist %s", id.String())
		}

		if !exists {
			listsExist = false
			msgs = append(msgs, fmt.Sprintf("blocklist %s does not exist", id.String()))
		}
	}

	// Validate session struct
	structValid, errStructValid := validator.ValidateStruct(s)
	if errStructValid != nil {
		msgs = append(msgs, errStructValid.Error())
	}

	var err error

	if len(msgs) > 0 {
		err = errors.New(strings.Join(msgs, ": "))
	}

	return listsExist && structValid, err
}

package session

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/jonathanfoster/freedom/store"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// Session represents a time frame in which websites are blocked
type Session struct {
	ID             uuid.UUID   `json:"id"`
	Name           string      `json:"name"`
	Starts         time.Time   `json:"starts" valid:"required"`
	Ends           time.Time   `json:"ends" valid:"required"`
	Blocklists     []uuid.UUID `json:"blocklists" valid:"required"`
	EverySunday    bool        `json:"every_sunday"`
	EveryMonday    bool        `json:"every_monday"`
	EveryTuesday   bool        `json:"every_tuesday"`
	EveryWednesday bool        `json:"every_wednesday"`
	EveryThursday  bool        `json:"every_thursday"`
	EveryFriday    bool        `json:"every_friday"`
	EverySaturday  bool        `json:"every_saturday"`
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
	weekday := now.Weekday()
	repeatsToday := false

	starts := s.Starts.UTC()
	ends := s.Ends.UTC()

	if weekday == time.Sunday && s.EverySunday {
		repeatsToday = true
	} else if weekday == time.Monday && s.EveryMonday {
		repeatsToday = true
	} else if weekday == time.Tuesday && s.EveryTuesday {
		repeatsToday = true
	} else if weekday == time.Wednesday && s.EveryWednesday {
		repeatsToday = true
	} else if weekday == time.Thursday && s.EveryThursday {
		repeatsToday = true
	} else if weekday == time.Friday && s.EveryFriday {
		repeatsToday = true
	} else if weekday == time.Saturday && s.EverySaturday {
		repeatsToday = true
	}

	if repeatsToday {
		starts = time.Date(now.Year(), now.Month(), now.Day(), starts.Hour(), starts.Minute(), starts.Second(), starts.Nanosecond(), starts.Location())
		ends = time.Date(now.Year(), now.Month(), now.Day(), ends.Hour(), ends.Minute(), ends.Second(), ends.Nanosecond(), ends.Location())
	}

	if starts.Before(now) && ends.After(now) {
		return true
	}

	return false
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
	return validator.ValidateStruct(s)
}

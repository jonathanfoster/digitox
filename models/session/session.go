package session

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/store"
)

// Session represents a time frame in which websites are blocked
type Session struct {
	ID             uuid.UUID              `json:"id" gorm:"type:text"`
	Name           string                 `json:"name"`
	Starts         time.Time              `json:"starts" valid:"required"`
	Ends           time.Time              `json:"ends" valid:"required"`
	Blocklists     []*blocklist.Blocklist `json:"blocklists" valid:"required" gorm:"many2many:session_blocklists"`
	EverySunday    bool                   `json:"every_sunday"`
	EveryMonday    bool                   `json:"every_monday"`
	EveryTuesday   bool                   `json:"every_tuesday"`
	EveryWednesday bool                   `json:"every_wednesday"`
	EveryThursday  bool                   `json:"every_thursday"`
	EveryFriday    bool                   `json:"every_friday"`
	EverySaturday  bool                   `json:"every_saturday"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	DeletedAt      *time.Time             `json:"deleted_at"`
}

// New creates a Session instance.
func New() *Session {
	return &Session{
		ID: uuid.NewV4(),
	}
}

// All retrieves all sessions from session store.
func All() ([]*Session, error) {
	var sessions []*Session

	if err := store.DB.Preload("Blocklists").Find(&sessions).Error; err != nil {
		return nil, errors.Wrap(err, "error retrieving all sessions")
	}

	return sessions, nil
}

// Exists checks if a session exists by ID.
func Exists(id uuid.UUID) (bool, error) {
	if _, err := Find(id); err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return false, nil
		}

		return false, errors.Wrap(err, "error checking if session exists")
	}

	return true, nil
}

// Find finds a session by ID.
func Find(id uuid.UUID) (*Session, error) {
	var sess Session

	if err := store.DB.Preload("Blocklists").Find(&sess, &Session{ID: id}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return nil, errors.Wrap(err, "error finding session")
	}

	return &sess, nil
}

// Remove removes the session from the file system.
func Remove(id uuid.UUID) error {
	exists, err := Exists(id)
	if err != nil {
		return errors.Wrap(err, "error removing session")
	}

	if !exists {
		return store.ErrNotFound
	}

	if err := store.DB.Delete(&Session{ID: id}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return errors.Wrap(err, "error removing session")
	}

	return nil
}

// IsActive determines whether a session is active based on starts, ends, and daily repeat options.
func (s *Session) IsActive() bool {
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

// Save writes the session to the session store.
func (s *Session) Save() error {
	if err := store.DB.Save(s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = store.ErrNotFound
		}

		return errors.Wrapf(err, "error saving session")
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (s *Session) Validate() (bool, error) {
	// TODO: Validate blocklist struct
	isValid, err := validator.ValidateStruct(s)
	if err != nil {
		return isValid, errors.Wrap(err, "error validating session")
	}

	return isValid, nil
}

package proxy

import (
	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
)

// Controller represents a structure responsible for controlling the state of
// the proxy blocklist in relation to active sessions.
type Controller struct{}

// NewController creates a Controller instance.
func NewController() *Controller {
	return &Controller{}
}

// RestartProxy restarts the proxy server so a new blocklist can take affect.
func (c *Controller) RestartProxy() error {
	// https://stackoverflow.com/a/30781156
	return nil
}

// Run starts a timer and updates proxy blocklist on a scheduled basis.
func (c *Controller) Run() error {
	// TODO: Loop on timer
	_, err := c.ActiveBlocklist()
	if err != nil {
		return err
	}

	restart := false

	// Compare expected blocklist to actual blocklist
	// Adjust actual blocklist
	// Domains missing? Add them.
	// Extra domains? Remove them.

	if restart {
		if err := c.RestartProxy(); err != nil {
			return err
		}
	}

	return nil
}

// ActiveBlocklist searches all sessions for active sessions and returns blocked domains.
func (c *Controller) ActiveBlocklist() ([]string, error) {
	var activeSessions []*session.Session

	// Get all sessions
	sessions, err := session.All()
	if err != nil {
		return nil, err
	}

	// Find active sessions
	for _, sess := range sessions {
		if sess.Active() {
			activeSessions = append(activeSessions, sess)
		}
	}

	var domains []string

	// Create active blocklist
	for _, sess := range activeSessions {
		for _, id := range sess.Blocklists {
			// Load blocklists
			list, err := blocklist.Find(id.String())
			if err != nil {
				return nil, err
			}

			// Copy blocked domains to active blocklist
			domains = append(domains, list.Domains...)
		}
	}

	return domains, nil
}

package proxy

import (
	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RestartProxy() error {
	return nil
}

func (c *Controller) Run() error {
	// TODO: Loop on timer
	restart, err := c.UpdateActiveBlocklist()
	if err != nil {
		return err
	}

	if restart {
		if err := c.RestartProxy(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Controller) UpdateActiveBlocklist() (bool, error) {
	var activeSessions []*session.Session
	var result = false

	// Get all sessions
	sessions, err := session.All()
	if err != nil {
		return false, err
	}

	// Find active sessions
	for _, sess := range sessions {
		if sess.Active() {
			activeSessions = append(activeSessions, sess)
		}
	}

	// Create expected blocklist
	activeList := blocklist.Blocklist{}
	for _, sess := range activeSessions {
		for _, list := range sess.Blocklists {
			// Load blocklists
			list, err = blocklist.Find(list.ID.String())
			if err != nil {
				return false, err
			}

			activeList.Domains = append(activeList.Domains, list.Domains...)
		}
	}

	// Compare expected blocklist to actual blocklist
	// Adjust actual blocklist
	// Domains missing? Add them.
	// Extra domains? Remove them.

	return result, nil
}

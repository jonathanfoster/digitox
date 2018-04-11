package proxy

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/store"
)

// Controller represents a structure responsible for controlling the state of
// the proxy blocklist in relation to active sessions.
type Controller struct {
	Filename   string
	Processing bool
	Tick       time.Duration
	Timeout    time.Duration
	ticker     *time.Ticker
}

// NewController creates a Controller instance.
func NewController(filename string) *Controller {
	return &Controller{
		Filename:   filename,
		Processing: false,
		Tick:       time.Second * 30,
		Timeout:    time.Second * 10,
	}
}

// ExpectedBlocklist returns the blocked domains from all active sessions.
func (c *Controller) ExpectedBlocklist() ([]string, error) {
	var activeSessions []*session.Session

	// Get all sessions
	sessions, err := session.All()
	if err != nil {
		if err == store.ErrNotExist {
			return nil, nil
		}

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
				// Be resilient to missing blocklists
				if err == store.ErrNotExist {
					log.Warnf("error determining expected blocklist: error finding blocklist %s: %s", id, err.Error())
					continue
				}

				return nil, err
			}

			// TODO: Remove duplicate domains
			// Copy blocked domains to active blocklist
			domains = append(domains, list.Domains...)
		}
	}

	return domains, nil
}

// ReadBlocklistFile retrieves the currently blocked domains from the proxy server.
func (c *Controller) ReadBlocklistFile() ([]string, error) {
	var list []string

	buf, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return list, nil
		}

		return nil, errors.Wrapf(err, "error opening blocklist file %s", c.Filename)
	}

	for _, s := range strings.Split(string(buf), "\n") {
		if s != "" {
			list = append(list, s)
		}
	}

	return list, nil
}

// RestartProxy restarts the proxy server so a new blocklist can take affect.
func (c *Controller) RestartProxy() error {
	// https://stackoverflow.com/a/30781156
	return nil
}

// Run starts a timer and updates proxy blocklist on a scheduled basis.
func (c *Controller) Run() {
	c.ticker = time.NewTicker(c.Tick)

	go func() {
		for range c.ticker.C {
			if c.Processing {
				log.Debug("proxy blocklist update processing: skipping tick")
				continue
			}

			c.Processing = true

			log.Debug("updating proxy blocklist")
			restart, err := c.UpdateBlocklist()
			if err != nil {
				log.Error("error updating blocklist in run loop: ", err)
			}

			if restart {
				log.Info("proxy blocklist updated: restarting proxy")
				if err := c.RestartProxy(); err != nil {
					log.Error("error restarting proxy run loop: ", err)
				}
			} else {
				log.Debug("proxy blocklist not updated: restart not required")
			}

			c.Processing = false
		}
	}()
}

// Stop stops controller and if processing, waits for the current process to end.
func (c *Controller) Stop() error {
	if c.ticker == nil {
		return nil
	}

	c.ticker.Stop()

	select {
	case <-time.After(time.Second * 10):
		return errors.New("proxy controller stop timeout expired")
	default:
		if !c.Processing {
			break
		}
	}

	return nil
}

// UpdateBlocklist updates proxy blocklist if changes are required.
func (c *Controller) UpdateBlocklist() (bool, error) {
	// Get expected blocklist from active sessions
	expected, err := c.ExpectedBlocklist()
	if err != nil {
		return false, err
	}

	// Get actual blocklist from proxy
	actual, err := c.ReadBlocklistFile()
	if err != nil {
		return false, err
	}

	// Compare expected blocklist to actual blocklist, update if not equal
	if !equals(expected, actual) {
		if err := c.WriteBlocklistFile(expected); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

// WriteBlocklistFile writes list to blocklist file.
func (c *Controller) WriteBlocklistFile(list []string) error {
	var buf bytes.Buffer

	for _, l := range list {
		if _, err := buf.WriteString(l + "\n"); err != nil {
			return errors.Wrap(err, "error writing blocklist string to buffer")
		}
	}

	if err := ioutil.WriteFile(c.Filename, buf.Bytes(), 0700); err != nil {
		return errors.Wrap(err, "error writing blocklist file")
	}

	return nil
}

func equals(expected []string, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}

	return true
}

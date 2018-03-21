package proxy

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Blocker represents blocking request middleware. If
// certain criteria is met, then it blocks the request
// returning 403 Forbidden. All other requests are passed
// along to the next HTTP request handler.
type Blocker struct {
	next http.Handler
}

// NewBlocker creates a new instance of Blocker.
func NewBlocker(next http.Handler) *Blocker {
	return &Blocker{
		next,
	}
}

func (b *Blocker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == "www.reddit.com" || r.Host == "www.reddit.com:443" {
		log.Warnf("%v %v %v %v", r.Method, r.URL, r.Proto, http.StatusForbidden)
		http.Error(w, fmt.Sprintf("host %s blocked", r.Host), http.StatusForbidden)
		return
	}

	b.next.ServeHTTP(w, r)
}

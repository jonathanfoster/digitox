package middleware

import (
	"net/http"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/server/handlers"
)

// Recovery represents a panic recovery middleware. When a panic occurs,
// the error is logged and an error response returned.
type Recovery struct{}

// NewRecovery creates a Recovery instance.
func NewRecovery() *Recovery {
	return &Recovery{}
}

func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			stack := make([]byte, 1024*8)
			stack = stack[:runtime.Stack(stack, true)]

			log.Errorf("recovered panic: %s: %s", err, stack)

			handlers.Error(w, http.StatusInternalServerError)
		}
	}()

	next(w, r)
}

package middleware

import (
	"net/http"
	"runtime"

	"github.com/jonathanfoster/freedom/api/httputil"
	"github.com/jonathanfoster/freedom/model"
	log "github.com/sirupsen/logrus"
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

			e := &model.Error{
				Message:    http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
			}

			httputil.JSON(w, http.StatusInternalServerError, e)
		}
	}()

	next(w, r)
}
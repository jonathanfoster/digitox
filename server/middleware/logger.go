package middleware

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// NewLogger creates a customized instance of negroni.Logger
// that uses logrus and doesn't prepend [negroni] to log entries.
func NewLogger() *negroni.Logger {
	l := negroni.NewLogger()
	l.ALogger = log.StandardLogger()
	l.SetFormat("{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}")

	return l
}

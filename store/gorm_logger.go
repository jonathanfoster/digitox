package store

import (
	log "github.com/sirupsen/logrus"
)

// GormLogger logs Gorm messages as debug messages.
type GormLogger struct{}

// Print logs Gorm messages as debug messages.
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.Debugf("gorm: sql: %s", v[3])
	case "log":
		log.Debugf("gorm: log: %s", v[2])
	}
}

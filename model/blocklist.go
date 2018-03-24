package model

import "github.com/satori/go.uuid"

// Blocklist represents a blocklist.
type Blocklist struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

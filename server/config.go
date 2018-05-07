package server

import (
	"crypto/rsa"
	"time"
)

// Config represents a server configuration
type Config struct {
	Addr              string
	ClientID          string
	ClientSecret      string
	DataSourceName    string
	TickerDuration    time.Duration
	TokenSigningKey   *rsa.PrivateKey
	TokenVerifyingKey *rsa.PublicKey
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	return &Config{}
}

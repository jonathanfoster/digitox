package server

import "crypto/rsa"

// Config represents a server configuration
type Config struct {
	Addr              string
	TokenSigningKey   *rsa.PrivateKey
	TokenVerifyingKey *rsa.PublicKey
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	return &Config{}
}

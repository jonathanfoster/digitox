package device

// Credentials represents a set of htpasswd credentials. This structure is used to keep the device struct decoupled
// from the htpasswd struct and avoid cyclical references.
type Credentials struct {
	dev *Device
}

// NewCredentials creates a Credentials instance.
func NewCredentials(dev *Device) *Credentials {
	return &Credentials{dev: dev}
}

// Username returns device name.
func (c *Credentials) Username() string {
	return c.dev.Name
}

// Password returns device password.
func (c *Credentials) Password() string {
	return c.dev.Password
}

// Hash returns device password hash.
func (c *Credentials) Hash() string {
	return c.dev.Hash
}
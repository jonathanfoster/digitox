package device

import (
	"github.com/jonathanfoster/digitox/store"
	"github.com/pkg/errors"
)

// Device represents a device.
type Device struct {
	Name     string `json:"name" valid:"required"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

// New creates a Device instance.
func New(name string) *Device {
	return &Device{
		Name: name,
	}
}

// All retrieves all devices.
func All() ([]*Device, error) {
	deviceNames, err := store.Device.All()
	if err != nil {
		return nil, err
	}

	devices := make([]*Device, len(deviceNames))

	for i, name := range deviceNames {
		devices[i] = New(name)
	}

	return devices, nil
}

// Find finds a blocklist by name.
func Find(name string) (*Device, error) {
	var hash string

	if err := store.Device.Find(name, &hash); err != nil {
		return nil, errors.Wrapf(err, "error finding device %s", name)
	}

	dev := New(name)
	dev.Hash = hash

	return dev, nil
}

// Save writes device to store.
func (d *Device) Save() error {
	credentials := NewCredentials(d)

	if err := store.Device.Save(d.Name, credentials); err != nil {
		return errors.Wrapf(err, "error saving device %s", d.Name)
	}

	return nil
}

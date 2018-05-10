package device

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/jonathanfoster/digitox/store"
	"github.com/pkg/errors"
)

// Device represents a device.
type Device struct {
	Name     string `json:"name" valid:"required"`
	Password string `json:"password" valid:"required"`
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

// Exists checks if a device exists by name.
func Exists(name string) (bool, error) {
	exists, err := store.Device.Exists(name)
	if err != nil {
		if err == store.ErrNotFound {
			return false, nil
		}

		return false, errors.Wrap(err, "error checking if device exists")
	}

	return exists, nil
}

// Find finds a blocklist by name.
func Find(name string) (*Device, error) {
	var dev Device

	if err := store.Device.Find(name, &dev); err != nil {
		return nil, errors.Wrap(err, "error finding device")
	}

	return &dev, nil
}

// Remove removes the device.
func Remove(name string) error {
	if err := store.Device.Remove(name); err != nil {
		return errors.Wrap(err, "error removing device")
	}

	return nil
}

// Save writes device to store.
func (d *Device) Save() error {
	if err := store.Device.Save(d.Name, d); err != nil {
		return errors.Wrap(err, "error saving device")
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (d *Device) Validate() (bool, error) {
	return validator.ValidateStruct(d)
}

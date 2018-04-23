package device

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/jonathanfoster/digitox/store"
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
		if err == store.ErrNotExist {
			return false, nil
		}

		return false, errors.Wrapf(err, "error checking if device %s exists", name)
	}

	return exists, nil
}

// Find finds a blocklist by name.
func Find(name string) (*Device, error) {
	var dev Device

	if err := store.Device.Find(name, &dev); err != nil {
		return nil, errors.Wrapf(err, "error finding device %s", name)
	}

	return &dev, nil
}

// Remove removes the device.
func Remove(name string) error {
	if err := store.Device.Remove(name); err != nil {
		return errors.Wrapf(err, "error removing device %s", name)
	}

	return nil
}

// Save writes device to store.
func (d *Device) Save() error {
	if err := store.Device.Save(d.Name, d); err != nil {
		return errors.Wrapf(err, "error saving device %s", d.Name)
	}

	return nil
}

// Validate validates tags for fields and returns false if there are any errors.
func (d *Device) Validate() (bool, error) {
	return validator.ValidateStruct(d)
}

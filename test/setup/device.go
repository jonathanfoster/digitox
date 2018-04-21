package setup

import (
	"os"
	"path"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/jonathanfoster/digitox/store"
)

var testDeviceFilename = os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/passwd"

// NewTestDevice creates a test device instance.
func NewTestDevice() *device.Device {
	dev := device.New("test")
	dev.Password = uuid.NewV4().String()

	return dev
}

// ResetTestDeviceStore removes all test devices and re-initializes the device store.
func ResetTestDeviceStore() {
	if err := os.RemoveAll(testDeviceFilename); err != nil {
		log.Panic("error resetting test device store: ", err.Error())
	}

	TestDeviceStore()
}

// TestDevice creates and saves a test device.
func TestDevice() *device.Device {
	dev := NewTestDevice()

	if err := dev.Save(); err != nil {
		log.Panic("error saving test device: ", err.Error())
	}

	return dev
}

// TestDeviceStore creates the test device directory and initializes the device store.
func TestDeviceStore() {
	dirname := path.Dir(testDeviceFilename)

	if err := os.MkdirAll(dirname, 0700); err != nil {
		log.Panicf("error creating test device directory %s: %s", dirname, err.Error())
	}

	store.Device = store.NewHtpasswdStore(testDeviceFilename)
}

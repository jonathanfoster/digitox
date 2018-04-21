package device_test

import (
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/jonathanfoster/digitox/test/setup"
)

func TestDevice(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Device", t, func() {
		setup.TestDeviceStore()
		testdev := setup.TestDevice()

		Convey("All", func() {
			Convey("Should return devices", func() {
				devices, err := device.All()
				So(err, ShouldBeNil)
				So(devices, ShouldNotBeEmpty)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				d := device.New("test")
				d.Password = uuid.NewV4().String()

				err := d.Save()
				So(err, ShouldBeNil)
			})
		})

		Convey("Find", func() {
			Convey("Should return device", func() {
				dev, err := device.Find(testdev.Name)
				So(err, ShouldBeNil)
				So(dev, ShouldNotBeNil)
			})
		})

		Reset(func() {
			setup.ResetTestDeviceStore()
		})
	})
}

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

		Convey("Find", func() {
			Convey("Should return device", func() {
				dev, err := device.Find(testdev.Name)
				So(err, ShouldBeNil)
				So(dev, ShouldNotBeNil)
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := device.Remove(testdev.Name)
				So(err, ShouldBeNil)
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

		Convey("Validate", func() {
			Convey("Should return true", func() {
				dev := setup.NewTestDevice()
				result, err := dev.Validate()
				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When password not provided", func() {
				Convey("Should return false", func() {
					dev := setup.NewTestDevice()
					dev.Password = ""
					result, err := dev.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})
		})

		Reset(func() {
			setup.ResetTestDeviceStore()
		})
	})
}

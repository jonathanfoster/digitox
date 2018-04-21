package device_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/jonathanfoster/digitox/test/setup"
)

func TestCredentials(t *testing.T) {
	Convey("Credentials", t, func() {
		dev := setup.NewTestDevice()
		credenitals := device.NewCredentials(dev)

		Convey("Username", func() {
			Convey("Should return device name", func() {
				So(credenitals.Username(), ShouldEqual, dev.Name)
			})
		})

		Convey("Password", func() {
			Convey("Should return device password", func() {
				So(credenitals.Password(), ShouldEqual, dev.Password)
			})
		})

		Convey("Hash", func() {
			Convey("Should return device hash", func() {
				So(credenitals.Hash(), ShouldEqual, dev.Hash)
			})
		})
	})
}

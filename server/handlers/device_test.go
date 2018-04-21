package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/jonathanfoster/digitox/test/setup"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/server"
)

func TestDeviceHandler(t *testing.T) {
	Convey("Device Handler", t, func() {
		router := server.NewRouter()
		setup.TestDeviceStore()
		testdev := setup.TestDevice()

		Convey("ListDevices", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/devices/", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 200)
			})
		})

		Convey("FindDevice", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/devices/"+testdev.Name, nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When device does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/devices/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("CreateDevice", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/devices/", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 501)
			})
		})

		Convey("UpdateDevice", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/devices/a8ae93e6-0e81-485e-9320-ff360fa70595", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 501)
			})
		})

		Convey("DeleteDevice", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/devices/a8ae93e6-0e81-485e-9320-ff360fa70595", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 501)
			})
		})

		Reset(func() {
			setup.ResetTestDeviceStore()
		})
	})
}

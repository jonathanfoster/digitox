package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/jonathanfoster/digitox/models/device"
	"github.com/jonathanfoster/digitox/test/setup"
	"github.com/satori/go.uuid"
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
			Convey("Status code should be 201", func() {
				dev := device.New(uuid.NewV4().String())
				dev.Password = uuid.NewV4().String()

				buf, err := json.Marshal(dev)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/devices/", buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 201)

				var body device.Device

				err = json.Unmarshal(w.Body.Bytes(), &body)
				So(err, ShouldBeNil)

				err = device.Remove(body.Name)
				So(err, ShouldBeNil)
			})

			Convey("When device is not valid", func() {
				Convey("Status code should be 422", func() {
					dev := device.New(uuid.NewV4().String())

					buf, err := json.Marshal(dev)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("POST", "/devices/", buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})
		})

		Convey("DeleteDevice", func() {
			Convey("Status code should be 204", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/devices/"+testdev.Name, nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 204)
			})

			Convey("When device does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/devices/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
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

		Reset(func() {
			setup.ResetTestDeviceStore()
		})
	})
}

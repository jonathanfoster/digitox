package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/server"
	"github.com/jonathanfoster/digitox/test/setup"
)

func TestSessionHandler(t *testing.T) {
	logrus.SetLevel(logrus.FatalLevel)

	Convey("Session Handler", t, func() {
		router := server.NewRouter()
		setup.TestBlocklistStore()
		setup.TestDeviceStore()
		setup.TestSessionStore()
		testlist := setup.TestBlocklist()
		testdev := setup.TestDevice()
		testsess := setup.TestSession(testlist.ID, testdev.Name)

		Convey("ListSessions", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions/", nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})
		})

		Convey("FindSession", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions/"+testsess.ID.String(), nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When session does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/sessions/doesnotexist", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("CreateSession", func() {
			Convey("Status code should be 201", func() {
				sess := setup.NewTestSession(testlist.ID, testdev.Name)
				sess.Name = "test"

				buf, err := json.Marshal(sess)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/sessions/", buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 201)

				var body session.Session

				err = json.Unmarshal(w.Body.Bytes(), &body)
				So(err, ShouldBeNil)

				err = session.Remove(body.ID.String())
				So(err, ShouldBeNil)
			})

			Convey("When session is not valid", func() {
				Convey("Status code should be 422", func() {
					sess := session.New() // ID must be a valid UUIDv4
					sess.Starts = time.Time{}

					buf, err := json.Marshal(sess)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("POST", "/sessions/", buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})
		})

		Convey("UpdateSession", func() {
			Convey("Status code should be 200", func() {
				testsess.Name = "test2"

				buf, err := json.Marshal(testsess)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/sessions/"+testsess.ID.String(), buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When session is not valid", func() {
				Convey("Status code should be 422", func() {
					origID := testsess.ID
					testsess.Starts = time.Time{}

					buf, err := json.Marshal(testsess)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/sessions/"+origID.String(), buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})

			Convey("When session does not exist", func() {
				Convey("Status code should be 404", func() {
					list := session.New()

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/sessions/"+list.ID.String(), buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("DeleteSession", func() {
			Convey("Status code should be 204", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/sessions/"+testsess.ID.String(), nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 204)
			})

			Convey("When session does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/sessions/doesnotexist", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID.String())
			session.Remove(testsess.ID.String())
			setup.ResetTestDeviceStore()
		})
	})
}

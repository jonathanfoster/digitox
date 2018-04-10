package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jonathanfoster/freedom/test/setup"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/server"
)

func TestSessionHandler(t *testing.T) {
	logrus.SetLevel(logrus.FatalLevel)

	Convey("Session Handler", t, func() {
		router := server.NewRouter()
		setup.TestSessionDirname()
		testsess := setup.TestSession()

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
					r := httptest.NewRequest("GET", "/sessions/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("CreateSession", func() {
			Convey("Status code should be 201", func() {
				sess := setup.NewTestSession()
				sess.Name = "test"

				buf, err := json.Marshal(sess)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/sessions/", buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 201)

				err = session.Remove(sess.ID.String())
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
					r := httptest.NewRequest("DELETE", "/sessions/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Reset(func() {
			session.Remove(testsess.ID.String())
		})
	})
}

package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/server"
	"github.com/jonathanfoster/freedom/test/testutil"
)

func TestSession(t *testing.T) {
	Convey("Session Handler", t, func() {
		router := server.NewRouter()

		if err := testutil.SetTestSessionDirname(); err != nil {
			panic(err)
		}

		testsess, err := testutil.CreateTestSession()
		if err != nil {
			panic(err)
		}

		Convey("ListSessions", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions", nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})
		})

		Convey("FindSession", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions/"+testsess.ID, nil)

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
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/sessions", nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 501)
			})
		})

		Convey("UpdateBlocklist", func() {
			Convey("Status code should be 200", func() {
				testsess.Blocklists = append(testsess.Blocklists, uuid.NewV4().String())

				buf, err := json.Marshal(testsess)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/sessions/"+testsess.ID, buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When session is not valid", func() {
				Convey("Status code should be 422", func() {
					origID := testsess.ID
					testsess.ID = "test" // ID must be a valid UUIDv4

					buf, err := json.Marshal(testsess)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/sessions/"+origID, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})

			Convey("When session does not exist", func() {
				Convey("Status code should be 404", func() {
					list := session.New(uuid.NewV4().String())

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/sessions/"+list.ID, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("DeleteSession", func() {
			Convey("Status code should be 204", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/sessions/"+testsess.ID, nil)

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
			session.Remove(testsess.ID)
		})
	})
}

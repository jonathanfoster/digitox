package session_test

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/test/testutil"
)

func TestSession(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Session", t, func() {
		if err := testutil.SetTestSessionDirname(); err != nil {
			panic(err)
		}

		testsess, err := testutil.SaveTestSession()
		if err != nil {
			panic(err)
		}

		Convey("All", func() {
			Convey("Should return sessions", func() {
				sess, err := session.All()
				So(err, ShouldBeNil)
				So(sess, ShouldNotBeEmpty)
			})
		})

		Convey("Find", func() {
			Convey("Should return session", func() {
				sess, err := session.Find(testsess.ID)
				So(err, ShouldBeNil)
				So(sess, ShouldNotBeEmpty)
			})

			Convey("Should load blocklists", func() {
				sess, err := session.Find(testsess.ID)
				So(err, ShouldBeNil)
				So(sess.Blocklists[0].ID, ShouldEqual, testsess.Blocklists[0].ID)
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := session.Remove(testsess.ID)
				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				sess := testutil.NewTestSession()
				sess.Name = "test-2"

				err := sess.Save()
				So(err, ShouldBeNil)

				err = session.Remove(sess.ID)
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Should return validation error", func() {
					testsess.ID = "test" // ID must be UUIDv4
					err := testsess.Save()
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("Validate", func() {
			Convey("Should return true", func() {
				sess := testutil.NewTestSession()
				result, err := sess.Validate()
				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When ID not a valid UUIDv4", func() {
				Convey("Should return false", func() {
					sess := testutil.NewTestSession()
					sess.ID = "test"
					result, err := sess.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When ID not provided", func() {
				Convey("Should return false", func() {
					sess := testutil.NewTestSession()
					sess.ID = ""
					result, err := sess.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When starts not provided", func() {
				Convey("Should return false", func() {
					sess := testutil.NewTestSession()
					sess.Starts = time.Time{}
					result, err := sess.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When ends not provided", func() {
				Convey("Should return false", func() {
					sess := testutil.NewTestSession()
					sess.Ends = time.Time{}
					result, err := sess.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When blocklists not provided", func() {
				Convey("Should return false", func() {
					sess := testutil.NewTestSession()
					sess.Blocklists = []blocklist.Blocklist{}
					result, err := sess.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})
		})

		Reset(func() {
			session.Remove(testsess.ID)
		})
	})
}

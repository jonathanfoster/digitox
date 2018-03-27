package session_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/model/session"
	"github.com/jonathanfoster/freedom/test/testutil"
)

func TestSession(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Session", t, func() {
		if err := testutil.SetTestSessionDirname(); err != nil {
			panic(err)
		}

		testsess, err := testutil.CreateTestSession()
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

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := session.Remove(testsess.ID)
				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				sess := session.New(uuid.NewV4().String())
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
				sess := session.New(uuid.NewV4().String())
				result, err := sess.Validate()
				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When ID not a valid UUIDv4", func() {
				Convey("Should return false", func() {
					sess := session.New("test")
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

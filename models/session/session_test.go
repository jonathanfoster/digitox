package session_test

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/test/setup"
)

func TestSession(t *testing.T) {
	log.SetLevel(log.FatalLevel)

	Convey("Session", t, func() {
		setup.TestSessionDirname()
		testsess := setup.TestSession()

		Convey("Active", func() {
			Convey("When session is active", func() {
				Convey("Should return true", func() {
					now := time.Now().UTC()
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 11, 59, 59, 0, now.Location())

					So(sess.Active(), ShouldBeTrue)
				})
			})

			Convey("When session is repeat and active", func() {
				Convey("Should return true", func() {
					now := time.Now().UTC().AddDate(0, 0, -1)
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 11, 59, 59, 0, now.Location())
					sess.EverySunday = true
					sess.EveryMonday = true
					sess.EveryTuesday = true
					sess.EveryWednesday = true
					sess.EveryThursday = true
					sess.EveryFriday = true
					sess.EverySaturday = true

					So(sess.Active(), ShouldBeTrue)
				})
			})

			Convey("When session is not active", func() {
				Convey("Should return false", func() {
					now := time.Now().UTC().AddDate(0, 0, -1)
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 11, 59, 59, 0, now.Location())

					So(sess.Active(), ShouldBeFalse)
				})
			})
		})

		Convey("All", func() {
			Convey("Should return sessions", func() {
				sess, err := session.All()

				So(err, ShouldBeNil)
				So(sess, ShouldNotBeEmpty)
			})
		})

		Convey("Find", func() {
			Convey("Should return session", func() {
				sess, err := session.Find(testsess.ID.String())

				So(err, ShouldBeNil)
				So(sess, ShouldNotBeEmpty)
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := session.Remove(testsess.ID.String())

				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				sess := setup.NewTestSession()
				sess.Name = "test-2"

				err := sess.Save()
				So(err, ShouldBeNil)

				err = session.Remove(sess.ID.String())
				So(err, ShouldBeNil)
			})
		})

		Convey("Validate", func() {
			Convey("Should return true", func() {
				sess := setup.NewTestSession()
				result, err := sess.Validate()

				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When starts not provided", func() {
				Convey("Should return false", func() {
					sess := setup.NewTestSession()
					sess.Starts = time.Time{}
					result, err := sess.Validate()

					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When ends not provided", func() {
				Convey("Should return false", func() {
					sess := setup.NewTestSession()
					sess.Ends = time.Time{}
					result, err := sess.Validate()

					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})

			Convey("When blocklists not provided", func() {
				Convey("Should return false", func() {
					sess := setup.NewTestSession()
					sess.Blocklists = nil
					result, err := sess.Validate()

					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})
		})

		Reset(func() {
			session.Remove(testsess.ID.String())
		})
	})
}

func TestSession_Active(t *testing.T) {
	Convey("Session", t, func() {

	})
}

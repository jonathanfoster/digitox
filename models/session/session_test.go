package session_test

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/test/setup"
)

func TestSession(t *testing.T) {
	log.SetLevel(log.FatalLevel)

	Convey("Session", t, func() {
		setup.TestBlocklistDirname()
		setup.TestSessionDirname()
		testlist := setup.TestBlocklist()
		testsess := setup.TestSessionWithBlocklist(testlist.ID)

		Convey("Active", func() {
			Convey("When session is active", func() {
				Convey("Should return true", func() {
					now := time.Now().UTC()
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

					So(sess.Active(), ShouldBeTrue)
				})
			})

			Convey("When session is repeat and active", func() {
				Convey("Should return true", func() {
					now := time.Now().UTC().AddDate(0, 0, -1)
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
					sess.RepeatEveryDay()

					So(sess.Active(), ShouldBeTrue)
				})
			})

			Convey("When session is not active", func() {
				Convey("Should return false", func() {
					now := time.Now().UTC().AddDate(0, 0, -1)
					sess := session.New()
					sess.Starts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
					sess.Ends = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

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

		Convey("Exists", func() {
			Convey("When session exists", func() {
				Convey("Should return true", func() {
					exists, err := session.Exists(testsess.ID.String())
					So(err, ShouldBeNil)
					So(exists, ShouldBeTrue)
				})
			})

			Convey("When session does not exist", func() {
				Convey("Should return false", func() {
					exists, err := session.Exists(uuid.NewV4().String())
					So(err, ShouldBeNil)
					So(exists, ShouldBeFalse)
				})
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

		Convey("RepeatEveryDay", func() {
			Convey("Should set every day to true", func() {
				testsess.RepeatEveryDay()

				So(testsess.EverySunday, ShouldBeTrue)
				So(testsess.EveryMonday, ShouldBeTrue)
				So(testsess.EveryTuesday, ShouldBeTrue)
				So(testsess.EveryWednesday, ShouldBeTrue)
				So(testsess.EveryThursday, ShouldBeTrue)
				So(testsess.EveryFriday, ShouldBeTrue)
				So(testsess.EverySaturday, ShouldBeTrue)
			})
		})

		Convey("RepeatNever", func() {
			Convey("Should set every day to false", func() {
				testsess.RepeatNever()

				So(testsess.EverySunday, ShouldBeFalse)
				So(testsess.EveryMonday, ShouldBeFalse)
				So(testsess.EveryTuesday, ShouldBeFalse)
				So(testsess.EveryWednesday, ShouldBeFalse)
				So(testsess.EveryThursday, ShouldBeFalse)
				So(testsess.EveryFriday, ShouldBeFalse)
				So(testsess.EverySaturday, ShouldBeFalse)
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
				valid, err := testsess.Validate()

				So(err, ShouldBeNil)
				So(valid, ShouldBeTrue)
			})

			Convey("When starts not provided", func() {
				Convey("Should return false", func() {
					testsess.Starts = time.Time{}
					valid, err := testsess.Validate()

					So(err, ShouldNotBeNil)
					So(valid, ShouldBeFalse)
				})
			})

			Convey("When ends not provided", func() {
				Convey("Should return false", func() {
					testsess.Ends = time.Time{}
					valid, err := testsess.Validate()

					So(err, ShouldNotBeNil)
					So(valid, ShouldBeFalse)
				})
			})

			Convey("When blocklists not provided", func() {
				Convey("Should return false", func() {
					testsess.Blocklists = nil
					valid, err := testsess.Validate()

					So(err, ShouldNotBeNil)
					So(valid, ShouldBeFalse)
				})
			})

			Convey("When blocklist does not exist", func() {
				Convey("Should return false", func() {
					testsess.Blocklists = append(testsess.Blocklists, uuid.NewV4())
					valid, err := testsess.Validate()

					So(err, ShouldNotBeNil)
					So(valid, ShouldBeFalse)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID.String())
			session.Remove(testsess.ID.String())
		})
	})
}

package session_test

import (
	"testing"

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

			Convey("Should load blocklists", func() {
				sess, err := session.Find(testsess.ID.String())
				So(err, ShouldBeNil)
				So(sess.Blocklists[0].ID, ShouldEqual, testsess.Blocklists[0].ID)
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

			//Convey("When sesssion is not valid", func() {
			//	Convey("Should return validation error", func() {
			//		testsess.ID = uuid.UUID{}
			//		err := testsess.Save()
			//		So(err, ShouldNotBeNil)
			//	})
			//})
		})

		//Convey("Validate", func() {
		//	Convey("Should return true", func() {
		//		sess := setup.NewTestSession()
		//		result, err := sess.Validate()
		//		So(err, ShouldBeNil)
		//		So(result, ShouldBeTrue)
		//	})
		//
		//	Convey("When ID not provided", func() {
		//		Convey("Should return false", func() {
		//			sess := setup.NewTestSession()
		//			sess.ID = uuid.UUID{}
		//			result, err := sess.Validate()
		//			So(err, ShouldNotBeNil)
		//			So(result, ShouldBeFalse)
		//		})
		//	})
		//
		//	Convey("When starts not provided", func() {
		//		Convey("Should return false", func() {
		//			sess := setup.NewTestSession()
		//			sess.Starts = time.Time{}
		//			result, err := sess.Validate()
		//			So(err, ShouldNotBeNil)
		//			So(result, ShouldBeFalse)
		//		})
		//	})
		//
		//	Convey("When ends not provided", func() {
		//		Convey("Should return false", func() {
		//			sess := setup.NewTestSession()
		//			sess.Ends = time.Time{}
		//			result, err := sess.Validate()
		//			So(err, ShouldNotBeNil)
		//			So(result, ShouldBeFalse)
		//		})
		//	})
		//
		//	Convey("When blocklists not provided", func() {
		//		Convey("Should return false", func() {
		//			sess := setup.NewTestSession()
		//			sess.Blocklists = []blocklist.Blocklist{}
		//			result, err := sess.Validate()
		//			So(err, ShouldNotBeNil)
		//			So(result, ShouldBeFalse)
		//		})
		//	})
		//})

		Reset(func() {
			session.Remove(testsess.ID.String())
		})
	})
}

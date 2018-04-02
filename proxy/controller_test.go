package proxy_test

import (
	"testing"
	"time"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/models/session"
	"github.com/jonathanfoster/freedom/proxy"
	"github.com/jonathanfoster/freedom/test/setup"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestController_ActiveBlocklist(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Controller", t, func() {
		setup.TestBlocklistDirname()
		setup.TestSessionDirname()
		testlist := setup.TestBlocklist()
		testsess := setup.TestSessionWithBlocklist(testlist.ID)

		Convey("ActiveBlocklist", func() {
			Convey("Should return active session blocklist domains", func() {
				ctrl := proxy.NewController()
				domains, err := ctrl.ActiveBlocklist()

				So(err, ShouldBeNil)
				So(domains, ShouldResemble, testlist.Domains)
			})

			Convey("When no active session", func() {
				Convey("Should return nil", func() {
					testsess.Starts = testsess.Starts.Add(time.Hour * 24)
					testsess.Ends = testsess.Ends.Add(time.Hour * 24)
					testsess.RepeatNever()
					So(testsess.Save(), ShouldBeNil)

					ctrl := proxy.NewController()
					domains, err := ctrl.ActiveBlocklist()

					So(err, ShouldBeNil)
					So(domains, ShouldBeNil)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID.String())
			session.Remove(testsess.ID.String())
		})
	})
}

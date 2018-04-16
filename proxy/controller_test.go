package proxy_test

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/proxy"
	"github.com/jonathanfoster/digitox/test/setup"
)

var filename = os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/active"

func TestController_ActiveBlocklist(t *testing.T) {
	Convey("Controller", t, func() {
		setup.TestBlocklistDirname()
		setup.TestSessionDirname()
		testlist := setup.TestBlocklist()
		testsess := setup.TestSession(testlist.ID)

		Convey("ActiveBlocklist", func() {
			Convey("Should return active session blocklist domains", func() {
				ctrl := proxy.NewController(filename)
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

					ctrl := proxy.NewController(filename)
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

func TestController_ReadBlocklistFile(t *testing.T) {
	Convey("Controller", t, func() {
		Convey("ReadBlocklistFile", func() {
			Convey("Should return blocklist", func() {
				list := setup.NewTestBlocklist()
				ctrl := proxy.NewController(filename)
				err := ctrl.WriteBlocklistFile(list.Domains)
				So(err, ShouldBeNil)

				domains, err := ctrl.ReadBlocklistFile()

				So(err, ShouldBeNil)
				So(domains, ShouldResemble, list.Domains)
			})

			Convey("When no blocklist", func() {
				Convey("Should return nil", func() {
					ctrl := proxy.NewController(filename)

					domains, err := ctrl.ReadBlocklistFile()

					So(err, ShouldBeNil)
					So(domains, ShouldBeNil)
				})
			})
		})

		Reset(func() {
			os.Remove(filename)
		})
	})
}

func TestController_UpdateBlocklist(t *testing.T) {
	Convey("Controller", t, func() {
		setup.TestBlocklistDirname()
		setup.TestSessionDirname()
		testlist := setup.TestBlocklist()
		testsess := setup.TestSession(testlist.ID)

		Convey("UpdateBlocklist", func() {
			Convey("When expected blocklist is not equal to actual blocklist", func() {
				Convey("Should update blocklist", func() {
					ctrl := proxy.NewController(filename)
					err := ctrl.WriteBlocklistFile(nil)
					So(err, ShouldBeNil)

					restart, err := ctrl.UpdateBlocklist()
					So(err, ShouldBeNil)
					So(restart, ShouldBeTrue)

					list, err := ctrl.ReadBlocklistFile()
					So(err, ShouldBeNil)
					So(list, ShouldResemble, testlist.Domains)
				})
			})

			Convey("When expected blocklist is equal to actual blocklist", func() {
				Convey("Should not update blocklist", func() {
					ctrl := proxy.NewController(filename)
					err := ctrl.WriteBlocklistFile(testlist.Domains)
					So(err, ShouldBeNil)

					restart, err := ctrl.UpdateBlocklist()
					So(err, ShouldBeNil)
					So(restart, ShouldBeFalse)
				})
			})

			Reset(func() {
				os.Remove(filename)
				blocklist.Remove(testlist.ID.String())
				session.Remove(testsess.ID.String())
			})
		})
	})
}

func TestController_WriteBlocklistFile(t *testing.T) {
	Convey("Controller", t, func() {
		Convey("WriteBlocklistFile", func() {
			Convey("Should not return error", func() {
				list := setup.NewTestBlocklist()
				ctrl := proxy.NewController(filename)
				err := ctrl.WriteBlocklistFile(list.Domains)

				So(err, ShouldBeNil)
			})

			Reset(func() {
				os.Remove(filename)
			})
		})
	})
}

package proxy_test

import (
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/session"
	"github.com/jonathanfoster/digitox/proxy"
	"github.com/jonathanfoster/digitox/test/setup"
)

var activeFilename = os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/active"

func TestController(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Proxy Controller", t, func() {
		setup.TestBlocklistStore()
		setup.TestSessionStore()
		testsess := setup.TestSession()

		var testDomains []string
		for _, domain := range testsess.Blocklists[0].Domains {
			testDomains = append(testDomains, domain)
		}

		Convey("ActiveBlocklist", func() {
			Convey("Should return active session blocklist domains", func() {
				ctrl := proxy.NewController(activeFilename)
				domains, err := ctrl.ActiveBlocklist()

				So(err, ShouldBeNil)
				So(domains, ShouldResemble, testDomains)
			})

			Convey("When no active session", func() {
				Convey("Should return nil", func() {
					testsess.Starts = testsess.Starts.Add(time.Hour * 24)
					testsess.Ends = testsess.Ends.Add(time.Hour * 24)
					testsess.RepeatNever()
					So(testsess.Save(), ShouldBeNil)

					ctrl := proxy.NewController(activeFilename)
					domains, err := ctrl.ActiveBlocklist()

					So(err, ShouldBeNil)
					So(domains, ShouldBeNil)
				})
			})
		})

		Convey("ReadBlocklistFile", func() {
			Convey("Should return blocklist", func() {
				list := setup.NewTestBlocklist()
				ctrl := proxy.NewController(activeFilename)
				err := ctrl.WriteBlocklistFile(list.Domains)
				So(err, ShouldBeNil)

				var expectedDomains []string
				for _, domain := range list.Domains {
					expectedDomains = append(expectedDomains, domain)
				}

				domains, err := ctrl.ReadBlocklistFile()

				So(err, ShouldBeNil)
				So(domains, ShouldResemble, expectedDomains)
			})

			Convey("When no blocklist", func() {
				Convey("Should return nil", func() {
					ctrl := proxy.NewController(activeFilename)

					domains, err := ctrl.ReadBlocklistFile()

					So(err, ShouldBeNil)
					So(domains, ShouldBeNil)
				})
			})
		})

		Convey("UpdateBlocklist", func() {
			Convey("When expected blocklist is not equal to actual blocklist", func() {
				Convey("Should update blocklist", func() {
					ctrl := proxy.NewController(activeFilename)
					err := ctrl.WriteBlocklistFile(nil)
					So(err, ShouldBeNil)

					restart, err := ctrl.UpdateBlocklist()
					So(err, ShouldBeNil)
					So(restart, ShouldBeTrue)

					list, err := ctrl.ReadBlocklistFile()
					So(err, ShouldBeNil)
					So(list, ShouldResemble, testDomains)
				})
			})

			Convey("When expected blocklist is equal to actual blocklist", func() {
				Convey("Should not update blocklist", func() {
					ctrl := proxy.NewController(activeFilename)
					err := ctrl.WriteBlocklistFile(testsess.Blocklists[0].Domains)
					So(err, ShouldBeNil)

					restart, err := ctrl.UpdateBlocklist()
					So(err, ShouldBeNil)
					So(restart, ShouldBeFalse)
				})
			})
		})

		Convey("WriteBlocklistFile", func() {
			Convey("Should not return error", func() {
				list := setup.NewTestBlocklist()
				ctrl := proxy.NewController(activeFilename)
				err := ctrl.WriteBlocklistFile(list.Domains)

				So(err, ShouldBeNil)
			})
		})

		Reset(func() {
			session.Remove(testsess.ID)
			os.Remove(activeFilename)
		})
	})
}

package blocklist_test

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/test/setup"
)

func TestBlocklist(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Blocklist", t, func() {
		setup.TestBlocklistDirname()
		testlist := setup.TestBlocklist()

		Convey("All", func() {
			Convey("Should return blocklists", func() {
				lists, err := blocklist.All()
				So(err, ShouldBeNil)
				So(lists, ShouldNotBeEmpty)
			})
		})

		Convey("Find", func() {
			Convey("Should return blocklist", func() {
				list, err := blocklist.Find(testlist.ID.String())
				So(err, ShouldBeNil)
				So(list, ShouldNotBeEmpty)
			})

			Convey("Should load hosts", func() {
				list, err := blocklist.Find(testlist.ID.String())
				So(err, ShouldBeNil)
				So(list.Hosts[0], ShouldEqual, testlist.Hosts[0])
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := blocklist.Remove(testlist.ID.String())
				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				list := blocklist.New()
				list.Hosts = append(list.Hosts, "www.reddit.com")

				err := list.Save()
				So(err, ShouldBeNil)

				err = blocklist.Remove(list.ID.String())
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Should return validation error", func() {
					testlist.Hosts = []string{}
					err := testlist.Save()
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("Validate", func() {
			Convey("Should return true", func() {
				list := setup.NewTestBlocklist()
				result, err := list.Validate()
				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When hosts are not provided", func() {
				Convey("Should return false", func() {
					list := setup.NewTestBlocklist()
					list.Hosts = []string{}
					result, err := list.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID.String())
		})
	})
}

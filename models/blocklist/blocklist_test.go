package blocklist_test

import (
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/models/blocklist"
	"github.com/jonathanfoster/digitox/test/setup"
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

		Convey("Exists", func() {
			Convey("When blocklist exists", func() {
				Convey("Should return true", func() {
					exists, err := blocklist.Exists(testlist.ID.String())
					So(err, ShouldBeNil)
					So(exists, ShouldBeTrue)
				})
			})

			Convey("When blocklist does not exist", func() {
				Convey("Should return false", func() {
					exists, err := blocklist.Exists(uuid.NewV4().String())
					So(err, ShouldBeNil)
					So(exists, ShouldBeFalse)
				})
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
				So(list.Domains[0], ShouldEqual, testlist.Domains[0])
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
				list.Domains = append(list.Domains, "www.reddit.com")

				err := list.Save()
				So(err, ShouldBeNil)

				err = blocklist.Remove(list.ID.String())
				So(err, ShouldBeNil)
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
					list.Domains = []string{}
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

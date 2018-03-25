package blocklist_test

import (
	"io/ioutil"
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/model/blocklist"
)

func TestBlocklist(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	Convey("Blocklist", t, func() {
		// Create a test blocklist before each test
		testlist := blocklist.New("test-" + uuid.NewV4().String())
		testlist.Hosts = append(testlist.Hosts, "www.reddit.com")
		if err := testlist.Save(); err != nil {
			panic(err)
		}

		Convey("All", func() {
			Convey("Should return blocklists", func() {
				lists, err := blocklist.All()
				So(err, ShouldBeNil)
				So(lists, ShouldNotBeEmpty)
			})
		})

		Convey("Find", func() {
			Convey("Should return blocklist", func() {
				list, err := blocklist.Find(testlist.Name)
				So(err, ShouldBeNil)
				So(list, ShouldNotBeEmpty)
			})

			Convey("Should load hosts", func() {
				list, err := blocklist.Find(testlist.Name)
				So(err, ShouldBeNil)
				So(list.Hosts[0], ShouldEqual, testlist.Hosts[0])
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := blocklist.Remove(testlist.Name)
				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				list := blocklist.New("test-" + uuid.NewV4().String())
				list.Hosts = append(list.Hosts, "www.reddit.com")

				err := list.Save()
				So(err, ShouldBeNil)

				err = blocklist.Remove(list.Name)
				So(err, ShouldBeNil)
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.Name)
		})
	})
}

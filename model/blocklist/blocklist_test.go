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
		testlist := blocklist.New(uuid.NewV4().String())
		testlist.Name = "test"
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
				list, err := blocklist.Find(testlist.ID)
				So(err, ShouldBeNil)
				So(list, ShouldNotBeEmpty)
			})

			Convey("Should load hosts", func() {
				list, err := blocklist.Find(testlist.ID)
				So(err, ShouldBeNil)
				So(list.Hosts[0], ShouldEqual, testlist.Hosts[0])
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				err := blocklist.Remove(testlist.ID)
				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				list := blocklist.New(uuid.NewV4().String())
				list.Hosts = append(list.Hosts, "www.reddit.com")

				err := list.Save()
				So(err, ShouldBeNil)

				err = blocklist.Remove(list.ID)
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Should return validation error", func() {
					testlist.ID = "test" // ID must be UUIDv4
					err := testlist.Save()
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("Validate", func() {
			Convey("Should return true", func() {
				list := blocklist.New(uuid.NewV4().String())
				result, err := list.Validate()
				So(err, ShouldBeNil)
				So(result, ShouldBeTrue)
			})

			Convey("When ID not a valid UUIDv4", func() {
				Convey("Should return false", func() {
					list := blocklist.New("test")
					result, err := list.Validate()
					So(err, ShouldNotBeNil)
					So(result, ShouldBeFalse)
				})
			})
		})

		Convey("Unmarshal", func() {
			Convey("Should unmarshal blocklist from data", func() {
				list := &blocklist.Blocklist{}
				data := []byte("# name: Social Distractions\n" +
					"www.reddit.com\n" +
					"news.ycombinator.com\n")

				err := blocklist.Unmarshal(data, list)
				So(err, ShouldBeNil)
				So(list.Name, ShouldEqual, "Social Distractions")
				So(list.Hosts[0], ShouldEqual, "www.reddit.com")
				So(list.Hosts[1], ShouldEqual, "news.ycombinator.com")
			})

			Convey("When name not provided", func() {
				Convey("Should unmarshal blocklist without name", func() {
					list := &blocklist.Blocklist{}
					data := []byte("www.reddit.com\n" +
						"news.ycombinator.com\n")

					err := blocklist.Unmarshal(data, list)
					So(err, ShouldBeNil)
					So(list.Name, ShouldEqual, "")
					So(list.Hosts[0], ShouldEqual, "www.reddit.com")
					So(list.Hosts[1], ShouldEqual, "news.ycombinator.com")
				})
			})

			Convey("When no lines provided", func() {
				Convey("Should unmarshal blocklist without name", func() {
					list := &blocklist.Blocklist{}
					data := []byte{}

					err := blocklist.Unmarshal(data, list)
					So(err, ShouldNotBeNil)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.Name)
		})
	})
}

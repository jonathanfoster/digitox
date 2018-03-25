package blocklist_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/model/blocklist"
)

func TestBlocklist(t *testing.T) {
	Convey("Blocklist", t, func() {
		Convey("All", func() {
			Convey("Should return blocklists", func() {
				lists, err := blocklist.All()

				So(err, ShouldBeNil)
				So(lists, ShouldNotBeEmpty)
			})
		})

		Convey("Find", func() {
			Convey("Should return blocklist", func() {
				list, err := blocklist.Find("default")

				So(err, ShouldBeNil)
				So(list, ShouldNotBeEmpty)
			})
		})

		Convey("Remove", func() {
			Convey("Should not return error", func() {
				name := "test-" + uuid.NewV4().String()
				err := ioutil.WriteFile(path.Join(blocklist.Dirname, name), nil, os.ModePerm)
				So(err, ShouldBeNil)

				err = blocklist.Remove(name)

				So(err, ShouldBeNil)
			})
		})

		Convey("Save", func() {
			Convey("Should not return error", func() {
				list := blocklist.New("test-" + uuid.NewV4().String())
				list.Hosts = append(list.Hosts, "www.reddit.com")

				err := list.Save()

				So(err, ShouldBeNil)

				if err := blocklist.Remove(list.Name); err != nil {
					So(err, ShouldBeNil)
				}
			})
		})
	})
}

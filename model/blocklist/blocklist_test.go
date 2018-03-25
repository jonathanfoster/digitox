package blocklist_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/model/blocklist"
)

func TestBlocklist(t *testing.T) {
	log.SetOutput(ioutil.Discard)

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

			Convey("Should load hosts", func() {
				host := "www.reddit.com"
				list := blocklist.New("test-" + uuid.NewV4().String())
				list.Hosts = append(list.Hosts, host)

				err := list.Save()
				So(err, ShouldBeNil)

				list, err = blocklist.Find(list.Name)
				So(err, ShouldBeNil)
				So(list.Hosts[0], ShouldEqual, host)

				err = blocklist.Remove(list.Name)
				So(err, ShouldBeNil)
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

				err = blocklist.Remove(list.Name)
				So(err, ShouldBeNil)
			})
		})
	})
}

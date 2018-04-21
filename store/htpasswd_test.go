package store_test

import (
	"os"
	"path"
	"testing"

	"github.com/jonathanfoster/digitox/store"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHtpasswdStore(t *testing.T) {
	Convey("Htpasswd Store", t, func() {
		Convey("Init", func() {
			Convey("When htpasswd file does not exist", func() {
				Convey("Should create htpasswd directory", func() {
					filename := os.Getenv("GOPATH") + "/src/github.com/jonathanfoster/digitox/bin/test/passwd"
					dirname := path.Dir(filename)

					os.Remove(filename)
					defer os.Remove(filename)

					_, err := os.Stat(filename)
					So(os.IsNotExist(err), ShouldBeTrue)

					h := store.NewHtpasswdStore(filename)
					h.Init()
					_, err = os.Stat(dirname)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

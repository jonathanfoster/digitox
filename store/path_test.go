package store_test

import (
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/store"
	"github.com/satori/go.uuid"
)

func TestPath(t *testing.T) {
	Convey("JoinPath", t, func() {
		dirname := "/etc/freedom/test/"

		Convey("Should join filename and directory", func() {
			id := uuid.NewV4().String()
			filename, err := store.JoinPath(id, dirname)
			So(err, ShouldBeNil)
			So(filename, ShouldEqual, path.Join(dirname, id))
		})

		Convey("Should sanitize filename", func() {
			id := "test=test"
			filename, err := store.JoinPath(id, dirname)
			So(err, ShouldBeNil)
			So(filename, ShouldEqual, path.Join(dirname, "test-test"))
		})

		Convey("When filename contains a relative path", func() {
			Convey("Should clean dot dot", func() {
				id := "../etc/passwd"
				filename, err := store.JoinPath(id, dirname)
				So(err, ShouldBeNil)
				So(filename, ShouldEqual, path.Join(dirname, "passwd"))
			})
		})
	})
}

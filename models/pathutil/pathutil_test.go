package pathutil_test

import (
	"path"
	"testing"

	"github.com/jonathanfoster/freedom/models/pathutil"
	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPathutil(t *testing.T) {
	Convey("Pathutil", t, func() {
		Convey("FileName", func() {
			dirname := "/etc/freedom/test/"

			Convey("Should join dirname directory and ID", func() {
				id := uuid.NewV4().String()
				filename, err := pathutil.FileName(id, dirname)
				So(err, ShouldBeNil)
				So(filename, ShouldEqual, path.Join(dirname, id))
			})

			Convey("Should sanitize ID", func() {
				id := "test=test"
				filename, err := pathutil.FileName(id, dirname)
				So(err, ShouldBeNil)
				So(filename, ShouldEqual, path.Join(dirname, "test-test"))
			})

			Convey("When ID contains a relative path", func() {
				Convey("Should clean dot dot", func() {
					id := "../etc/passwd"
					filename, err := pathutil.FileName(id, dirname)
					So(err, ShouldBeNil)
					So(filename, ShouldEqual, path.Join(dirname, "passwd"))
				})
			})
		})
	})
}

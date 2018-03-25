package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/api"
	"github.com/jonathanfoster/freedom/model/blocklist"
)

func TestBlocklist(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	Convey("Blocklist Handler", t, func() {
		router := api.NewRouter()

		Convey("ListBlocklists", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/blocklists", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 200)
			})
		})

		Convey("FindBlocklist", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/blocklists/default", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 200)
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/blocklists/notfound", nil)

					router.ServeHTTP(w, r)

					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("DeleteBlocklist", func() {
			Convey("Status code should be 204", func() {
				name := "test-" + uuid.NewV4().String()
				err := ioutil.WriteFile(path.Join(blocklist.Dirname, name), nil, os.ModePerm)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/blocklists/"+name, nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 204)
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/blocklists/notfound", nil)

					router.ServeHTTP(w, r)

					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("CreateBlocklist", func() {
			Convey("Status code should be 201", func() {
				list := blocklist.New("test-" + uuid.NewV4().String())
				list.Hosts = append(list.Hosts, "www.reddit.com")

				buf, err := json.Marshal(list)
				buffer := bytes.NewBuffer(buf)

				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/blocklists", buffer)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 201)
			})
		})
	})
}

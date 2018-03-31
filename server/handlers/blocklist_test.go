package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/server"
	"github.com/jonathanfoster/freedom/test/setup"
)

func TestBlocklistHandler(t *testing.T) {
	log.SetLevel(log.FatalLevel)

	Convey("Blocklist Handler", t, func() {
		router := server.NewRouter()
		setup.TestBlocklistDirname()
		testlist := setup.TestBlocklist()

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
				r := httptest.NewRequest("GET", "/blocklists/"+testlist.ID.String(), nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/blocklists/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("DeleteBlocklist", func() {
			Convey("Status code should be 204", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/blocklists/"+testlist.ID.String(), nil)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 204)
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/blocklists/notfound", nil)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("CreateBlocklist", func() {
			Convey("Status code should be 201", func() {
				list := blocklist.New()
				list.Name = "test"
				list.Domains = append(list.Domains, "www.reddit.com")

				buf, err := json.Marshal(list)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/blocklists", buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 201)

				err = blocklist.Remove(list.ID.String())
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					list := setup.NewTestBlocklist()
					list.Domains = []string{}

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("POST", "/blocklists", buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})
		})

		Convey("UpdateBlocklist", func() {
			Convey("Status code should be 200", func() {
				testlist.Domains = append(testlist.Domains, "news.ycombinator.com")

				buf, err := json.Marshal(testlist)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/blocklists/"+testlist.ID.String(), buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					testlist.Domains = []string{}

					buf, err := json.Marshal(testlist)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+testlist.ID.String(), buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 404", func() {
					list := blocklist.New()

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+list.ID.String(), buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID.String())
		})
	})
}

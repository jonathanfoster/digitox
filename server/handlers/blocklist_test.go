package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/models/blocklist"
	"github.com/jonathanfoster/freedom/server"
	"github.com/jonathanfoster/freedom/test/testutil"
)

func TestBlocklist(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	Convey("Blocklist Handler", t, func() {
		router := server.NewRouter()

		if err := testutil.SetTestBlocklistDirname(); err != nil {
			panic(err)
		}

		testlist, err := testutil.SaveTestBlocklist()
		if err != nil {
			panic(err)
		}

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
				r := httptest.NewRequest("GET", "/blocklists/"+testlist.ID, nil)

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
				r := httptest.NewRequest("DELETE", "/blocklists/"+testlist.ID, nil)

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
				list := blocklist.New(uuid.NewV4().String())
				list.Name = "test"
				list.Hosts = append(list.Hosts, "www.reddit.com")

				buf, err := json.Marshal(list)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/blocklists", buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 201)

				err = blocklist.Remove(list.ID)
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					list := blocklist.New("test") // ID must be a valid UUIDv4
					list.Name = "test"
					list.Hosts = append(list.Hosts, "www.reddit.com")

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
				testlist.Hosts = append(testlist.Hosts, "news.ycombinator.com")

				buf, err := json.Marshal(testlist)
				buffer := bytes.NewBuffer(buf)
				So(err, ShouldBeNil)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/blocklists/"+testlist.ID, buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					origID := testlist.ID
					testlist.ID = "test" // ID must be a valid UUIDv4

					buf, err := json.Marshal(testlist)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+origID, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 404", func() {
					list := blocklist.New(uuid.NewV4().String())

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+list.ID, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.ID)
		})
	})
}

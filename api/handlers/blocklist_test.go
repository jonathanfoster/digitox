package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
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

		// Create a test blocklist before each test
		testlist := blocklist.New("test-" + uuid.NewV4().String())
		testlist.Hosts = append(testlist.Hosts, "www.reddit.com")
		if err := testlist.Save(); err != nil {
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
				r := httptest.NewRequest("GET", "/blocklists/"+testlist.Name, nil)

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
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/blocklists/"+testlist.Name, nil)

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

				err = blocklist.Remove(list.Name)
				So(err, ShouldBeNil)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					list := blocklist.New("") // Name must not be empty
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
				r := httptest.NewRequest("PUT", "/blocklists/"+testlist.Name, buffer)

				router.ServeHTTP(w, r)
				So(w.Code, ShouldEqual, 200)
			})

			Convey("When blocklist is not valid", func() {
				Convey("Status code should be 422", func() {
					testlist.Dirname = "" // Dirname must not be empty

					buf, err := json.Marshal(testlist)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+testlist.Name, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 422)
				})
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 404", func() {
					list := blocklist.New("test-" + uuid.NewV4().String())

					buf, err := json.Marshal(list)
					buffer := bytes.NewBuffer(buf)
					So(err, ShouldBeNil)

					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/blocklists/"+list.Name, buffer)

					router.ServeHTTP(w, r)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Reset(func() {
			blocklist.Remove(testlist.Name)
		})
	})
}

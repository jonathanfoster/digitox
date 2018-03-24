package handlers_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/jonathanfoster/freedom/api"
)

func TestBlocklist(t *testing.T) {
	Convey("Blocklist Handler", t, func() {
		router := api.NewRouter()

		Convey("ListBlocklists", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/blocklists", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 200, w.Code)
			})
		})

		Convey("FindBlocklist", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/blocklists/default", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 200, w.Code)
			})

			Convey("When blocklist does not exist", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/blocklists/notfound", nil)

					router.ServeHTTP(w, r)

					assert.Equal(t, 404, w.Code)
				})
			})
		})
	})
}

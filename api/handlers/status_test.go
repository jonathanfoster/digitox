package handlers_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/api"
)

func TestStatus(t *testing.T) {
	Convey("Status Handler", t, func() {
		router := api.NewRouter()

		Convey("Status", func() {
			Convey("Status code should be 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 200)
			})
		})
	})
}

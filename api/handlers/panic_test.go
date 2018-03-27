package handlers_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/freedom/api"
)

func TestPanic(t *testing.T) {
	Convey("Panic Handler", t, func() {
		router := api.NewRouter()

		Convey("Panic", func() {
			Convey("Status code should be 500", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/panic", nil)

				router.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 500)
			})
		})
	})
}

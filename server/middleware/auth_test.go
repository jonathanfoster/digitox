package middleware_test

import (
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/negroni"

	"github.com/jonathanfoster/digitox/server"
	"github.com/jonathanfoster/digitox/server/middleware"
	"github.com/jonathanfoster/digitox/server/oauth"
)

func TestAuth(t *testing.T) {
	log.SetLevel(log.FatalLevel)

	Convey("Auth", t, func() {
		n := negroni.New()
		n.Use(middleware.NewAuth(oauth.DefaultVerifyingKey))
		n.UseHandler(server.NewRouter())

		Convey("Root path (/)", func() {
			Convey("Should not require authorization", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/", nil)

				n.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 200)
			})
		})

		Convey("OAuth token path (/oauth/token)", func() {
			Convey("Should not require authorization", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/oauth/token", nil)

				n.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 400)
			})
		})

		Convey("All other paths", func() {
			Convey("Should require authorization", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/blocklists/", nil)

				n.ServeHTTP(w, r)

				So(w.Code, ShouldEqual, 401)
			})
		})
	})
}

package handlers_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/jonathanfoster/freedom/api"
)

func TestSession(t *testing.T) {
	Convey("Session Handler", t, func() {
		router := api.NewRouter()

		Convey("ListSessions", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 501, w.Code)
			})
		})

		Convey("FindSession", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/sessions/a8ae93e6-0e81-485e-9320-ff360fa70595", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 501, w.Code)
			})

			Convey("When ID is not a UUID", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/sessions/1234567890", nil)

					router.ServeHTTP(w, r)

					assert.Equal(t, 400, w.Code)
				})
			})
		})

		Convey("CreateSession", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/sessions", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 501, w.Code)
			})
		})

		Convey("UpdateSession", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/sessions/a8ae93e6-0e81-485e-9320-ff360fa70595", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 501, w.Code)
			})

			Convey("When ID is not a UUID", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("PUT", "/sessions/1234567890", nil)

					router.ServeHTTP(w, r)

					assert.Equal(t, 400, w.Code)
				})
			})
		})

		Convey("DeleteSession", func() {
			Convey("Status code should be 501", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/sessions/a8ae93e6-0e81-485e-9320-ff360fa70595", nil)

				router.ServeHTTP(w, r)

				assert.Equal(t, 501, w.Code)
			})

			Convey("When ID is not a UUID", func() {
				Convey("Status code should be 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/sessions/1234567890", nil)

					router.ServeHTTP(w, r)

					assert.Equal(t, 400, w.Code)
				})
			})
		})
	})
}

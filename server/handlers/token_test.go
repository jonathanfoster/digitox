package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jonathanfoster/digitox/server"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func TestOAuthHandler(t *testing.T) {
	log.SetLevel(log.FatalLevel)

	Convey("Blocklist Handler", t, func() {
		router := server.NewRouter()
		clientID := "admin"
		clientSecret := "Digitox123"

		Convey("Token", func() {
			Convey("Client Credentials grant type", func() {
				Convey("GET /oauth/token", func() {
					targetFormat := "/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s&redirect_uri=http://localhost"
					target := fmt.Sprintf(targetFormat, clientID, clientSecret)

					Convey("Status code should be 200", func() {
						w := httptest.NewRecorder()
						r := httptest.NewRequest("GET", target, nil)

						router.ServeHTTP(w, r)
						So(w.Code, ShouldEqual, 200)
					})

					Convey("Should return access token", func() {
						w := httptest.NewRecorder()
						r := httptest.NewRequest("GET", target, nil)

						router.ServeHTTP(w, r)
						So(w.Code, ShouldEqual, 200)

						var token AccessToken
						err := json.Unmarshal(w.Body.Bytes(), &token)
						So(err, ShouldBeNil)

						So(token.AccessToken, ShouldNotBeEmpty)
						So(token.TokenType, ShouldEqual, "Bearer")
						So(token.ExpiresIn, ShouldNotBeEmpty)
					})

					Convey("When client is unauthorized", func() {
						Convey("Status code should be 400", func() {
							w := httptest.NewRecorder()
							r := httptest.NewRequest("GET", fmt.Sprintf(targetFormat, "doesnotexist", "doesnotexist"), nil)

							router.ServeHTTP(w, r)
							So(w.Code, ShouldEqual, 400)
						})
					})
				})
			})
		})
	})
}

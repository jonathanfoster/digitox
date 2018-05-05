package middleware

import (
	"crypto/rsa"
	"net/http"
	"regexp"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/server/handlers"
)

// All paths require authorization except the following
var noAuthRegexp = []*regexp.Regexp{
	regexp.MustCompile(`^\/oauth\/.*$`), // /oauth/token
	regexp.MustCompile(`^\/$`),          // /
}

// Auth represents authorization middleware.
type Auth struct {
	jwtMiddleware *jwtmiddleware.JWTMiddleware
}

// NewAuth creates a Auth instance.
func NewAuth(verifyingKey *rsa.PublicKey) *Auth {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			log.Error("error handling auth token: ", err)
			handlers.Error(w, http.StatusUnauthorized)
		},
		Extractor: jwtmiddleware.FromFirst(jwtmiddleware.FromAuthHeader, jwtmiddleware.FromParameter("access_token")),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return verifyingKey, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
		UserProperty:  "cid",
	})

	return &Auth{
		jwtMiddleware: jwtMiddleware,
	}
}

// All paths require authorization unless matched with no auth list
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	for _, re := range noAuthRegexp {
		if re.MatchString(r.URL.Path) {
			next(w, r)
			return
		}
	}

	a.jwtMiddleware.HandlerWithNext(w, r, next)
}

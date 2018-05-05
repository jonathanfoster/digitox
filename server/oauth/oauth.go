package oauth

import (
	"crypto/rsa"

	"github.com/RangelReale/osin"

	"github.com/jonathanfoster/digitox/store"
)

const (
	DefaultClientID     = "59f92849-b883-402c-b429-15a67663d4f3"
	DefaultClientSecret = "450a31ea-0c18-4925-97db-b9f981ca4a62"
)

// Server serves OAuth endpoint requests.
var Server *osin.Server

func init() {
	// Initialize OAuth server with default signing key for testing purposes
	InitOAuthServer(DefaultSigningKey, DefaultClientID, DefaultClientSecret)
}

// InitOAuthServer initializes the OAuth server.
func InitOAuthServer(signingKey *rsa.PrivateKey, clientID string, clientSecret string) {
	config := osin.NewServerConfig()
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.CLIENT_CREDENTIALS}
	config.AllowGetAccessRequest = true
	config.AllowClientSecretInParams = true
	config.ErrorStatusCode = 400

	Server = osin.NewServer(config, store.NewOAuthStore(clientID, clientSecret))
	Server.AccessTokenGen = NewJWTAccessTokenGen(signingKey)
}

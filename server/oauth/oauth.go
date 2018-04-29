package oauth

import (
	"crypto/rsa"

	"github.com/RangelReale/osin"

	"github.com/jonathanfoster/digitox/store"
)

// Server serves OAuth endpoint requests.
var Server *osin.Server

func init() {
	// Initialize OAuth server with default signing key for testing purposes
	InitOAuth(DefaultSigningKey)
}

// InitOAuth initializes the OAuth server.
func InitOAuth(signingKey *rsa.PrivateKey) {
	config := osin.NewServerConfig()
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.CLIENT_CREDENTIALS}
	config.AllowGetAccessRequest = true
	config.AllowClientSecretInParams = true
	config.ErrorStatusCode = 400

	Server = osin.NewServer(config, store.NewOAuthStore())
	Server.AccessTokenGen = NewJWTAccessTokenGen(signingKey)
}

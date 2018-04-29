package oauth

import (
	"crypto/rsa"

	"github.com/RangelReale/osin"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// JWTAccessTokenGen represents a JWT access token generator.
type JWTAccessTokenGen struct {
	signingKey *rsa.PrivateKey
}

// NewJWTAccessTokenGen creates a JWTAccessTokenGen instance
func NewJWTAccessTokenGen(signingKey *rsa.PrivateKey) *JWTAccessTokenGen {
	return &JWTAccessTokenGen{
		signingKey: signingKey,
	}
}

// GenerateAccessToken generates a JWT access token.
func (j *JWTAccessTokenGen) GenerateAccessToken(data *osin.AccessData, generateRefresh bool) (accessToken string, refreshToken string, err error) {
	signingMethod := jwt.SigningMethodRS256

	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accessToken, err = token.SignedString(j.signingKey)
	if err != nil {
		return "", "", errors.Wrap(err, "error signing access token")
	}

	if !generateRefresh {
		return
	}

	token = jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		"cid": data.Client.GetId(),
	})

	refreshToken, err = token.SignedString(j.signingKey)
	if err != nil {
		return "", "", errors.Wrap(err, "error signing refresh token")
	}

	return
}

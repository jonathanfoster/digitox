package handlers

import (
	"net/http"

	"github.com/RangelReale/osin"
	log "github.com/sirupsen/logrus"

	"github.com/jonathanfoster/digitox/server/oauth"
)

// Token handles the GET, POST /oauth/token route.
func Token(w http.ResponseWriter, r *http.Request) {
	res := oauth.Server.NewResponse()
	defer res.Close()

	if ar := oauth.Server.HandleAccessRequest(res, r); ar != nil {
		switch ar.Type {
		case osin.CLIENT_CREDENTIALS, osin.REFRESH_TOKEN:
			ar.Authorized = true
		}

		oauth.Server.FinishAccessRequest(res, r, ar)
	}

	if res.IsError {
		if res.InternalError != nil {
			log.Error("error handling access request: ", res.InternalError.Error())
		} else {
			log.Warn("access request failed: ", res.ErrorId)
		}
	}

	if err := osin.OutputJSON(res, w, r); err != nil {
		log.Error("error outputting json: ", err.Error())
	}
}

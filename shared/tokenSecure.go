package shared

import (
	"log"
	"net/http"

	"github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
)

// ValidateRequest will verify that a token received from an http request
// is valid and signy by authority
func ValidateRequest(domain, audience string, req *http.Request) bool {

	var auth0Domain = "https://" + domain + "/"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
	configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	_, err := validator.ValidateRequest(req)

	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

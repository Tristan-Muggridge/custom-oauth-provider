package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Move to a shared package
type AppRegistration struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	FirstName    string
	LastName     string
	Email        string
}

type AppRegistrationsRepository interface {
	Find(clientId string) (AppRegistration, error)
}

type AppRegistrationsRepositoryImpl struct {
	Registrations map[string]AppRegistration
}

func (repo AppRegistrationsRepositoryImpl) Find(clientId string) (AppRegistration, error) {
	appRegistration, ok := repo.Registrations[clientId]
	if !ok {
		return AppRegistration{}, errors.New("not found")
	}
	return appRegistration, nil
}

// End of shared package

var appRegistrationsRepository AppRegistrationsRepository = AppRegistrationsRepositoryImpl{
	Registrations: map[string]AppRegistration{
		"123": {
			ClientId:     "123",
			ClientSecret: "secret",
			RedirectUri:  "http://localhost:8080/api/v1/hello-world",
			FirstName:    "Tristan",
			LastName:     "Muggridge",
			Email:        "muggridge.dev@gmail.com",
		},
	},
}

type AuthorizeParams struct {
	ResponseType string
	ClientId     string
	RedirectUri  string
	Scope        string
	State        string
}

var AuthController = struct {
	Authorize func(c *gin.Context)
}{
	Authorize: func(c *gin.Context) {
		// TODO: Grab the response_type, client_id, redirect_uri, scope, and state from the request

		authorizeParams := AuthorizeParams{
			ResponseType: c.Query("response_type"),
			ClientId:     c.Query("client_id"),
			RedirectUri:  c.Query("redirect_uri"),
			Scope:        c.Query("scope"),
			State:        c.Query("state"),
		}

		// TODO: Check valid clientId
		appRegistration, err := appRegistrationsRepository.Find(authorizeParams.ClientId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid_client",
			})
			return
		}

		// TODO: Check valid redirectURI

		if authorizeParams.RedirectUri != appRegistration.RedirectUri {
			c.JSON(400, gin.H{
				"error": "invalid_redirect_uri",
			})
			return
		}

		c.Redirect(http.StatusFound, authorizeParams.RedirectUri)
	},
}

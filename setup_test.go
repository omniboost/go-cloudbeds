package cloudbeds_test

import (
	"os"

	"github.com/omniboost/go-cloudbeds"
	"golang.org/x/oauth2"
)

func client() *cloudbeds.Client {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	refreshToken := os.Getenv("OAUTH_REFRESH_TOKEN")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	oauthConfig := cloudbeds.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client := cloudbeds.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	return client
}

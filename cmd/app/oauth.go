package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// initGoogleOauth2Config initializes the Google OAuth2 configuration.
// To get Google OAuth2 credentials: https://console.cloud.google.com/apis/credentials
func initGoogleOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  googleOauth2CallbackURL,
		ClientID:     googleOauth2ClientID,
		ClientSecret: googleOauth2ClientSecret,
		Scopes:       googleOauth2Scopes,
		Endpoint:     google.Endpoint,
	}
}

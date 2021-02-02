package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  = "random-test"
)

func InitAuth() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/google-callback",
		ClientID:     Env("GOOGLE_CLIENT_ID", ""),
		ClientSecret: Env("GOOGLE_CLIENT_SECRET", ""),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

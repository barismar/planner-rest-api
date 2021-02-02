package services

import "golang.org/x/oauth2"

var (
	GoogleOauthConfig *oauth2.Config
	oauthStateString  = "random-test"
)

func handleGoogleLogin() string {
	url := GoogleOauthConfig.AuthCodeURL(oauthStateString)

	return url
}

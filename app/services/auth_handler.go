package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barismar/planner-rest-api/config"
	"golang.org/x/oauth2"
)

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_at"`
	Picture       string `json:"picture"`
}

func GetUserInfo(state string, code string) (UserInfo, error) {
	userInfo := UserInfo{}
	if state != config.OauthStateString {
		return userInfo, fmt.Errorf("invalid oauth state")
	}

	token, err := config.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return userInfo, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?oauth_token=" + token.AccessToken)
	if err != nil {
		return userInfo, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return userInfo, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	jsonErr := json.Unmarshal(contents, &userInfo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return userInfo, nil
}

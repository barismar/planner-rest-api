package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/barismar/planner-rest-api/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func Login(context *gin.Context) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.OauthStateString)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(context *gin.Context) {
	log.SetFormatter(&log.JSONFormatter{})
	content, err := getUserInfo(
		context.Query("state"),
		context.Query("code"),
	)
	if err != nil {
		fmt.Println(err.Error())
		context.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": content,
	})
}

type Profile struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_at"`
	Picture       string `json:"picture"`
}

func getUserInfo(state string, code string) (Profile, error) {
	profile1 := Profile{}
	if state != config.OauthStateString {
		return profile1, fmt.Errorf("invalid oauth state")
	}

	token, err := config.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return profile1, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?oauth_token=" + token.AccessToken)
	if err != nil {
		return profile1, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return profile1, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	jsonErr := json.Unmarshal(contents, &profile1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Println(profile1)

	return profile1, nil
}

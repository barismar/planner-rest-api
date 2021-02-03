package controllers

import (
	"fmt"
	"net/http"

	"github.com/barismar/planner-rest-api/app/services"
	"github.com/barismar/planner-rest-api/config"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.OauthStateString)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(context *gin.Context) {
	content, err := services.GetUserInfo(
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

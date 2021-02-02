package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	homeHtml := []byte(`<html><body><a href="/login">login</a></body></html>`)
	context.Data(
		http.StatusOK,
		"text/html; charset=utf-8",
		homeHtml,
	)
}

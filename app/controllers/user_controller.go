package controllers

import (
	"net/http"

	"github.com/barismar/planner-rest-api/app/models"
	"github.com/gin-gonic/gin"
)

func StoreUser(context *gin.Context) {
	var input models.CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"data":    models.StoreUser(input),
	})
}

func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    models.GetUsers(),
	})
}

func ShowUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    models.ShowUser(context.Param("id")),
	})
}

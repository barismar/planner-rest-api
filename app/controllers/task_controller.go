package controllers

import (
	"net/http"

	"github.com/barismar/planner-rest-api/app/models"
	"github.com/gin-gonic/gin"
)

func StoreTask(context *gin.Context) {
	var input models.CreateTaskInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"data":    models.StoreTask(input),
	})
}

func GetTasks(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    models.GetTasks(),
	})
}

func ShowTask(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    models.ShowTask(context.Param("id")),
	})
}

package controllers

import (
	"net/http"

	"github.com/barismar/planner-rest-api/app/models"
	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	Description string `json:"description" binding:"required"`
}

func StoreTask(context *gin.Context) {
	var input CreateTaskInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"data":    models.StoreTask(input.Description),
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

package controllers

import (
	"net/http"

	"github.com/barismar/planner-rest-api/app/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    models.GetTasks(),
	})
}

func ShowTask(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    context.Param("id"),
	})
}

package routes

import (
	"github.com/barismar/planner-rest-api/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	ApiRoute(router)

	return router
}

func ApiRoute(route *gin.Engine) {
	route.GET("/health-check", controllers.HealthCheck)

	route.GET("/tasks", controllers.GetTasks)
	route.GET("/tasks/:id", controllers.ShowTask)
}

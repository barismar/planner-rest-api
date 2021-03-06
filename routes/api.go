package routes

import (
	"github.com/barismar/planner-rest-api/app/controllers"
	"github.com/barismar/planner-rest-api/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	ApiRoute(router)
	TemplateRoute(router)
	AuthRoute(router)

	return router
}

func ApiRoute(route *gin.Engine) {
	route.GET("/health-check", controllers.HealthCheck)

	route.POST("/users", controllers.StoreUser)
	route.GET("/users", controllers.GetUsers)
	route.GET("/users/:id", controllers.ShowUser)

	route.POST("/tasks", controllers.StoreTask)
	route.GET("/tasks", controllers.GetTasks)
	route.GET("/tasks/:id", controllers.ShowTask)

}

func TemplateRoute(route *gin.Engine) {
	route.GET("/", controllers.Home)
}

func AuthRoute(route *gin.Engine) {
	config.InitAuth()
	route.GET("/login", controllers.Login)
	route.GET("/google-callback", controllers.GoogleCallback)
}

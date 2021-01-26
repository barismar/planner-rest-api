package main

import (
	"log"

	"github.com/barismar/planner-rest-api/app/models"
	"github.com/barismar/planner-rest-api/config"
	"github.com/barismar/planner-rest-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	models.DB = config.InitDatabase()
	router := routes.SetupRouter()

	router.Run(":" + config.Env("APP_PORT", "3000"))
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

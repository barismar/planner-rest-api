package main

import (
	"log"
	"os"

	"github.com/barismar/planner-rest-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	router := routes.SetupRouter()

	router.Run(":" + os.Getenv("APP_PORT"))
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

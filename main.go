package main

import (
	"github.com/barismar/planner-rest-api/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run(":3000")
}

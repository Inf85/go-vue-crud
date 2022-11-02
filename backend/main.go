package main

import (
	"go-vue-grud/backend/database"
	"go-vue-grud/backend/helpers"
	"go-vue-grud/backend/routes"
	"os"
)

func main() {
	logger := helpers.GetLogger()

	dbErr := database.InitDB()
	if dbErr != nil {
		logger.Error(dbErr.Error())
		os.Exit(1)
	}

	logger.Println("Starting the HTTP server on port 9080")
	routes.Routes()

}

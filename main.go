package main

import (
	"github.com/davidklz/ds-app-backend/database"
	"github.com/davidklz/ds-app-backend/pkgs/logger"
	"github.com/davidklz/ds-app-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("The .env file could not be loaded: %s", err.Error())
	}

	store, err := database.ConnectToPostgres()
	if err != nil {
		logger.Fatal("Error while trying to create postgres store:\n %s", err.Error())
	}

	server := server.InitializeServer(":3000", store)
	server.Run()
}

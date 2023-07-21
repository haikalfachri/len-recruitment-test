package main

import (
	"len-test/database"
	"len-test/routes"
	"len-test/utils"

	"github.com/labstack/echo"
)

func main() {
	config := database.Config{
		DB_USERNAME: utils.GetConfig("POSTGRES_USER"),
		DB_PASSWORD: utils.GetConfig("POSTGRES_PASSWORD"),
		DB_HOST:     utils.GetConfig("POSTGRES_HOST"),
		DB_PORT:     utils.GetConfig("POSTGRES_PORT"),
		DB_NAME:     utils.GetConfig("POSTGRES_DB"),
	}

	config.ConnectDB()

	database.MigrateDB()

	e := echo.New()
	routes.SetUpRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
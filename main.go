package main

import (
	"mpesa/database"
	"mpesa/migration"
	"mpesa/router"
)

func main() {
	//Initialising db ...
	database.InitDb()
	// Migrate the schema
	migration.Migrate()
	//Creating routers
	router.InitServer()
}

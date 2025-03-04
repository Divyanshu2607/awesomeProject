package main

import (
	"API/database"
	"API/routes"
)

func main() {
	database.InitDB() // Initialize database
	r := routes.SetupRouter()
	r.Run(":8080") // Start server
}

package main

import (
	"backend/database"
	"backend/routes"
)

func main() {
	// Setup database connection
	database.SetupDatabase()

	// Inisialisasi Router dari routes
	r := routes.SetupRouter()

	// Run server
	r.Run(":8080") // Bisa ganti port jika diperlukan
}

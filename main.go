package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Setup database connection
	database.SetupDatabase()

	// Inisialisasi Router dari routes
	r := routes.SetupRouter()

	// Ambil PORT dari environment variable (untuk deploy online)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default ke 8080 jika tidak ada PORT di env
	}

	// Jalankan server
	fmt.Println("Server running on port:", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

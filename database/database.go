package database

import (
	"api-web-pemerintah/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func SetupDatabase() {
	// koneksi ke PostgreSQL di Railway
	dsn := "postgres://postgres:wsKcDFfPnBqqEwfLhyERIlYGNIjmRJse@junction.proxy.rlwy.net:22215/railway?sslmode=disable"
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connected successfully!")

	fmt.Println("Running database migration...")
	DB.AutoMigrate(&models.Users{}, &models.Aspirations{}, models.Admins{})
	fmt.Println("Migration done.")
}

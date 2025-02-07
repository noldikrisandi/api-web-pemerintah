package database

import (
	"backend/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

// SetupDatabase initializes the database connection
func SetupDatabase() {
	// Setup koneksi ke PostgreSQL
	dsn := "user=postgres password=noldi dbname=webpemerintah sslmode=disable"
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connected successfully!")

	// Migrate model: pastikan tabel News dan Informasi ada di database
	fmt.Println("Running database migration...")
	DB.AutoMigrate(&models.Users{}, &models.Informations{}, &models.News{}, &models.Subsidies{}, &models.Aspirations{}, &models.Destinations{}, &models.Testimonies{}, &models.Packages{}, &models.Visitors{}, &models.Orders{})
	fmt.Println("Migration done.")
}

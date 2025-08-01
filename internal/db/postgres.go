package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/adityjoshi/parkingSystem.git/internal/models"
)

var DB *gorm.DB

func Init() {
	dsn := "host=postgres user=user password=password dbname=parkingdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Connected to DB")

	DB.AutoMigrate(&models.Vehicle{}, &models.ParkingSlot{}, &models.ParkingSession{})
}

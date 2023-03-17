package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")                                  // Loading the DB config from .env
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Connects to the DB

	if err != nil {
		panic("Failed to connect to DB!")
	}
}

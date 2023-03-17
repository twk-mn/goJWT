package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load() // Loading .env

	if err != nil {
		log.Fatal("Error laoding .env file!")
	}
}

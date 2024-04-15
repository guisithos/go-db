package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnviroment() {

	{
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		serviceName := os.Getenv("DB_SERVICE_NAME")

	}
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	go_ora "github.com/sijms/go-ora/v2"
)

func ConnectToDatabase() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	serviceName := os.Getenv("DB_SERVICE_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connStr := go_ora.BuildUrl(host, port, serviceName, user, password, nil)

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to Oracle database successful")

	return db, nil
}

func CloseConnectionDb() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	serviceName := os.Getenv("DB_SERVICE_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connStr := go_ora.BuildUrl(host, port, serviceName, user, password, nil)

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	db.Close()

	fmt.Println("Connection to Oracle database closed.")
	return db, nil
}
